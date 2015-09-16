package contracts

import (
	"einvite/framework"
	"time"
)

type User struct {
	Email string
	Name  string
}

type UserAuthCredentials struct {
	Type         framework.AuthType
	AccessToken  string
	RefreshToken string
	Expiry       time.Time
}

type SessionInfo struct {
	Id     string
	Values map[string]string
	User   *framework.SessionUser
	Expiry time.Time
}
