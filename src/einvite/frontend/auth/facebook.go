package auth

import (
	"code.google.com/p/goauth2/oauth"
	"einvite/common/contracts"
	"einvite/common/services"
	"einvite/framework"
	"encoding/json"
	"fmt"
	"log"
)

type FacebookController struct {
	profileInfoURL string
	oauthCfg       *oauth.Config
	userService    services.UserService
}

// Start the authorization process
func (this *FacebookController) Auth(ctx framework.WebContext) framework.WebResult {

	//can be used to pass state between server-to-server calls (roundtripped)
	state := ""

	//Get the Google URL which shows the Authentication page to the user
	url := this.oauthCfg.AuthCodeURL(state)

	//redirect user to that page
	return ctx.Redirect(url)
}

// Function that handles the callback from the Google server
func (this *FacebookController) AuthCallback(ctx framework.WebContext) framework.WebResult {
	//Get the code from the response
	code, _ := ctx.Param("code")

	t := &oauth.Transport{Config: this.oauthCfg}

	// Exchange the received code for a token
	t.Exchange(code)

	//now get user data based on the Transport which has the token
	resp, err := t.Client().Get(this.profileInfoURL)

	if err != nil {
		log.Println(err)
		return ctx.FrameworkError(framework.ToError(framework.Error_Web_UnableToAuthenticate, err))
	}

	userprofile := make(map[string]interface{})
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&userprofile)

	if err != nil {
		log.Println(err)
		return ctx.FrameworkError(framework.ToError(framework.Error_Web_UnableToAuthenticate, err))
	}

	fmt.Println(userprofile)

	email := userprofile["email"].(string)
	name := userprofile["name"].(string)

	userDto := &contracts.User{Email: email, Name: name}
	userCredentialsDto := &contracts.UserAuthCredentials{
		Type:         framework.AuthType_Facebook,
		AccessToken:  t.Token.AccessToken,
		RefreshToken: t.Token.RefreshToken,
		Expiry:       t.Token.Expiry,
	}

	user, err2 := this.userService.SaveWithCredentials(userDto, userCredentialsDto)

	if err2 != nil {
		return ctx.Error(err2)
	}

	ctx.Session().SetUser(&framework.SessionUser{
		UserId:   user.Email,
		AuthType: framework.AuthType_Facebook,
		AuthData: t.AccessToken,
	})

	return ctx.Template(userInfoTemplate, fmt.Sprintf("Name: %s Email: %s", name, email))
}

func NewFacebookController(userService services.UserService) *FacebookController {

	cfg := &authConfig{}
	framework.Config.ReadInto("facebook", &cfg)

	oauthCfg := &oauth.Config{
		ClientId:       cfg.ClientId,
		ClientSecret:   cfg.ClientSecret,
		AuthURL:        cfg.AuthURL,
		TokenURL:       cfg.TokenURL,
		RedirectURL:    cfg.RedirectURL,
		Scope:          cfg.Scope,
		ApprovalPrompt: cfg.ApprovalPrompt,
		AccessType:     cfg.AccessType,
	}

	return &FacebookController{profileInfoURL: cfg.ProfileInfoURL, oauthCfg: oauthCfg, userService: userService}
}
