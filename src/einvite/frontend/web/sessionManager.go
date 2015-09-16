package web

import (
	"einvite/common/contracts"
	"einvite/common/services"
	"einvite/framework"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type webSessionManager struct {
	config         *WebConfig
	sessionService services.SessionService
}

func (this *webSessionManager) Get(ctx *webContext) (*webSession, error) {

	//read cookies
	sessionCookie, _ := ctx.Request.Cookie(this.config.Session.Cookie)
	if sessionCookie != nil {

		log.Println("Session cookie found")
		return this.tryLoadSession(ctx, sessionCookie.Value)
	}

	//TODO: support url sessions?
	//log.Println("New session")
	return this.newSession(ctx), nil
}

func (this *webSessionManager) Clear(ctx *webContext) {

	sessionCookie := &http.Cookie{}
	sessionCookie.Name = this.config.Session.Cookie
	sessionCookie.MaxAge = -1
	sessionCookie.Expires = time.Unix(1, 0)

	log.Println("Clearing session cookie")

	ctx.SetCookie(sessionCookie)
}

func (this *webSessionManager) Save(ctx *webContext) error {

	var sessionId string
	var err error

	sessionCfg := this.config.Session
	securityCfg := this.config.Security

	duration := time.Duration(sessionCfg.Ttl) * time.Second

	ctx.session.expiry = time.Now().UTC().Add(duration)

	if ctx.session.hasChanges() {

		info := &contracts.SessionInfo{
			Id:     ctx.session.id,
			Values: ctx.session.values,
			User:   ctx.session.user,
			Expiry: ctx.session.expiry,
		}

		sessionId, err = this.sessionService.Save(info)

		if err != nil {
			return err
		}

		if ctx.session.IsNew() {
			signed, encrypted := framework.Security.EncryptAndSign(sessionId, securityCfg.RawSignKey, securityCfg.RawEncryptionKey)
			publicId := fmt.Sprintf("%s#%s", signed, encrypted)
			ctx.session.setId(sessionId, publicId)
		}
	} else {

		if sessionCfg.AutoRefresh {
			//just refresh the session
			//log.Println(fmt.Sprintf("Setting session %s expiry to %s", ctx.session.id, ctx.session.expiry.String()))
			this.sessionService.SetExpiry(ctx.session.id, ctx.session.expiry)
		}
	}

	sessionCookie := &http.Cookie{}
	sessionCookie.Name = sessionCfg.Cookie
	sessionCookie.MaxAge = sessionCfg.Ttl
	sessionCookie.Expires = ctx.session.expiry
	sessionCookie.HttpOnly = sessionCfg.HttpOnly
	sessionCookie.Secure = sessionCfg.Secure
	sessionCookie.Value = ctx.session.publicId
	sessionCookie.Path = "/"
	//log.Println("Setting cookie " + sessionCookie.String())
	ctx.SetCookie(sessionCookie)

	return nil
}

func (this *webSessionManager) tryLoadSession(ctx *webContext, publicId string) (*webSession, error) {

	securityConfig := this.config.Security

	separator := strings.Index(publicId, "#")
	if separator < 0 {
		return nil, framework.NewError(framework.Error_Web_SessionTampered, "session signature mismatch")
	}

	signature := publicId[:separator]
	encrypted := publicId[separator+1:]

	//verify signature
	ok := framework.Security.VerifySignature(signature, encrypted, securityConfig.RawSignKey)
	if !ok {
		return nil, framework.NewError(framework.Error_Web_SessionTampered, "session signature mismatch")
	}

	//its safe to decrypt
	sessionId := framework.Security.Decrypt(encrypted, securityConfig.RawEncryptionKey)

	info, err := this.sessionService.Get(sessionId)
	//could have been expired?
	if err != nil {
		return nil, err
	}

	if info == nil {
		//something bad happened
		return nil, framework.NewError(framework.Error_Web_SessionNotFound, "Session not found")
	}

	fmt.Println(fmt.Sprintf("Loading session with publicId %s and id %s", publicId, sessionId))
	session := this.loadSession(ctx, publicId, info)
	return session, nil
}

func (this *webSessionManager) newSession(context *webContext) *webSession {

	session := &webSession{}

	session.id = ""
	session.values = make(map[string]string)

	session.changed = false

	return session
}

func (this *webSessionManager) loadSession(context *webContext, publicId string, info *contracts.SessionInfo) *webSession {

	session := &webSession{}
	session.publicId = publicId //Sign(Encrypt(id)) + # + Encrypt(id)
	session.id = info.Id

	session.expiry = info.Expiry

	if info.Values != nil {
		session.values = info.Values
	} else {
		session.values = make(map[string]string)
	}

	if info.User != nil {
		session.user = &framework.SessionUser{
			UserId:   info.User.UserId,
			AuthType: info.User.AuthType,
			AuthData: info.User.AuthData,
		}
	}

	session.changed = false

	return session
}

func NewSessionManager(config *WebConfig, sessionService services.SessionService) *webSessionManager {

	return &webSessionManager{config: config, sessionService: sessionService}
}
