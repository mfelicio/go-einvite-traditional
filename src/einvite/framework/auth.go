package framework

type AuthType int

const (
	AuthType_Google   AuthType = 1
	AuthType_Facebook AuthType = 2
	AuthType_Twitter  AuthType = 3
)

type googleAuthData struct {
}
