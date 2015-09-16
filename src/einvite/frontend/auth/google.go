package auth

import (
	"code.google.com/p/goauth2/oauth"
	"einvite/common/contracts"
	"einvite/common/services"
	"einvite/framework"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"time"
)

var notAuthenticatedTemplate = template.Must(template.New("").Parse(`
<html><body>
You aren't logged in. Please authenticate this app with one of the OAuth providers below.
<form action="/auth/google" method="POST"><input type="submit" value="Login with Google account"/></form>
<form action="/auth/facebook" method="POST"><input type="submit" value="Login with Facebook account"/></form>
</body></html>
`))

var userInfoTemplate = template.Must(template.New("").Parse(`
<html><body>
This app is now authenticated to access your Google user info.  Your details are:<br />
{{.}}
</body></html>
`))

var authenticatedTemplate = template.Must(template.New("").Parse(`
<html><body>
Hello:<br />
<b>{{.}}</b>
</body></html>
`))

type GoogleController struct {
	profileInfoURL string
	oauthCfg       *oauth.Config
	userService    services.UserService
}

var randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))

func (this *GoogleController) HandleRoot(ctx framework.WebContext) framework.WebResult {

	session := ctx.Session()

	if sessionUser := session.User(); sessionUser == nil {

		return ctx.Template(notAuthenticatedTemplate, nil)
	} else {
		email := sessionUser.UserId
		user, _ := this.userService.Get(email)
		return ctx.Template(authenticatedTemplate, user.Name)
	}
}

// Start the authorization process
func (this *GoogleController) Auth(ctx framework.WebContext) framework.WebResult {

	//can be used to pass state between server-to-server calls (roundtripped)
	state := ""

	//Get the Google URL which shows the Authentication page to the user
	url := this.oauthCfg.AuthCodeURL(state)

	//redirect user to that page
	return ctx.Redirect(url)
}

// Function that handles the callback from the Google server
func (this *GoogleController) AuthCallback(ctx framework.WebContext) framework.WebResult {
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

	email := userprofile["email"].(string)
	name := userprofile["name"].(string)

	userDto := &contracts.User{Email: email, Name: name}
	userCredentialsDto := &contracts.UserAuthCredentials{
		Type:         framework.AuthType_Google,
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
		AuthType: framework.AuthType_Google,
		AuthData: t.AccessToken,
	})

	return ctx.Template(userInfoTemplate, fmt.Sprintf("Name: %s Email: %s", name, email))
}

func (this *GoogleController) TestHighLoad(ctx framework.WebContext) framework.WebResult {

	n := randomizer.Int63()
	email := fmt.Sprintf("user%d@einvite.com", n)
	name := fmt.Sprintf("User %d", n)

	dto := &contracts.User{Email: email, Name: name}
	credentials := &contracts.UserAuthCredentials{
		Type:         framework.AuthType_Google,
		AccessToken:  fmt.Sprintf("Access token for user %s", name),
		RefreshToken: fmt.Sprintf("Refresh token for user %s", name),
		Expiry:       time.Now().Add(1 * time.Hour),
	}
	_, err := this.userService.SaveWithCredentials(dto, credentials)
	if err != nil {
		log.Println(err)
		return ctx.Error(err)
	}

	session := ctx.Session()

	session.SetUser(&framework.SessionUser{
		UserId:   email,
		AuthType: framework.AuthType_Google,
		AuthData: credentials.AccessToken,
	})

	return ctx.Text("OK")
}

func NewGoogleController(userService services.UserService) *GoogleController {

	cfg := &authConfig{}
	framework.Config.ReadInto("google", &cfg)

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

	return &GoogleController{profileInfoURL: cfg.ProfileInfoURL, oauthCfg: oauthCfg, userService: userService}
}
