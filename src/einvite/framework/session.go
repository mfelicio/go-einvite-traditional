package framework

import (
	"time"
)

type Session interface {
	IsNew() bool

	User() *SessionUser
	SetUser(user *SessionUser) error

	Remove(name string) bool
	Get(name string) (string, bool)
	//Returns <oldValue,true> or <newValue,false>
	//Bool flag means Updated
	Set(name string, value string) (string, bool)

	Expiry() time.Time

	Save()
}

type SessionUser struct {
	UserId   string
	AuthType AuthType
	AuthData interface{}
}
