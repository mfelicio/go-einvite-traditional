package web

import (
	"einvite/common/services"
	"einvite/framework"
	"fmt"
	_web "github.com/hoisie/web"
	"io/ioutil"
	"log"
)

type webServer struct {
	config         *WebConfig
	sessionManager *webSessionManager
	_srv           *_web.Server
}

func (this *webServer) ListenAndServe() {
	log.Println("Starting http server on port", this.config.Http.Port)
	this._srv.Config.StaticDir = this.config.Http.StaticPath
	this._srv.Run("0.0.0.0:" + this.config.Http.Port)
}

func (this *webServer) Get(route string, fn framework.WebHandler) {

	this._srv.Get(route, func(_ctx *_web.Context) {
		this.processWebHandler(_ctx, fn)
	})
}

func (this *webServer) Post(route string, fn framework.WebHandler) {

	this._srv.Post(route, func(_ctx *_web.Context) {
		this.processWebHandler(_ctx, fn)
	})
}

func (this *webServer) Put(route string, fn framework.WebHandler) {
	this._srv.Put(route, func(_ctx *_web.Context) {
		this.processWebHandler(_ctx, fn)
	})
}

func (this *webServer) Delete(route string, fn framework.WebHandler) {
	this._srv.Delete(route, func(_ctx *_web.Context) {
		this.processWebHandler(_ctx, fn)
	})
}

func (this *webServer) processWebHandler(_ctx *_web.Context, fn framework.WebHandler) {
	context := NewContext(this, _ctx)

	session, err := this.sessionManager.Get(context)

	if err == nil {
		//set the session
		context.session = session

		//process the request
		result := fn(context)

		if result == nil {
			result = &GenericResult{}
		}

		//process response header
		if session.User() != nil {
			//log.Println("Saving session")
			this.sessionManager.Save(context)
		}

		//process response body
		err = result.Write(context)
		if err != nil {
			//TODO: may not work because some data may have been writen
			this.processError(context, err)
		}

	} else {
		log.Println("Error reading session")
		this.sessionManager.Clear(context)
		this.processError(context, err)
	}

}

func (this *webServer) processError(context *webContext, err error) {
	var fErr *framework.FrameworkError

	switch err.(type) {

	case *framework.FrameworkError:
		fErr = err.(*framework.FrameworkError)
	default:
		fErr = framework.ToError(framework.Error_Generic, err)
	}

	log.Println(fmt.Sprintf("Aborting due to %s", fErr.ErrorMessage))
	context.Abort(400, fErr.ErrorMessage)
}

func NewWebServer(webCfg *WebConfig, sessionService services.SessionService) framework.WebServer {

	sessionManager := NewSessionManager(webCfg, sessionService)
	_srv := _web.NewServer()
	_srv.Logger = log.New(ioutil.Discard, log.Prefix(), log.Flags())
	server := &webServer{config: webCfg, sessionManager: sessionManager, _srv: _srv}

	return server
}
