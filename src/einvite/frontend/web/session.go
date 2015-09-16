package web

import (
	"einvite/framework"
	"time"
)

type webSession struct {
	id       string
	publicId string

	expiry time.Time

	user *framework.SessionUser

	values map[string]string

	changed bool
}

func (this *webSession) User() *framework.SessionUser {
	return this.user
}

func (this *webSession) SetUser(user *framework.SessionUser) error {
	if this.user == nil {
		this.user = user
		this.changedIf(true)
		return nil
	}

	return framework.NewError(framework.Error_Web_SessionAlreadyHasUser, "session already has user")
}

func (this *webSession) IsNew() bool {
	return this.id == ""
}

func (this *webSession) Expiry() time.Time {
	return this.expiry
}

func (this *webSession) Get(name string) (string, bool) {
	value, ok := this.values[name]
	return value, ok
}

func (this *webSession) Set(name string, value string) (string, bool) {
	old, updated := this.values[name]
	this.values[name] = value

	if updated {
		this.changedIf(old != value)
		return old, true
	} else {

		this.changedIf(true)
		return value, false
	}
}

func (this *webSession) Remove(name string) bool {
	var ok bool
	if _, ok = this.values[name]; ok {
		delete(this.values, name)
	}

	return ok
}

func (this *webSession) Save() {
	this.changedIf(true)
}

func (this *webSession) changedIf(value bool) {
	if !this.changed {
		this.changed = value
	}
}

func (this *webSession) hasChanges() bool {
	return this.IsNew() || this.changed
}

func (this *webSession) setId(sessionId string, publicId string) error {
	if this.IsNew() {

		this.id = sessionId
		this.publicId = publicId
		this.changedIf(true)

		return nil
	} else {
		return framework.NewError(framework.Error_Web_SessionAlreadyHasId, "session already has id")
	}

}
