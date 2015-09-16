package framework

import ()

type WebResult interface {
	Write(ctx WebContext) error
}
