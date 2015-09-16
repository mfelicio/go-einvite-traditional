package web

import (
	"bytes"
	"einvite/framework"
	_web "github.com/hoisie/web"
	"html/template"
	"net/http"
	"time"
)

type webContext struct {
	server        *webServer
	*_web.Context //embedded

	session *webSession
}

func (this *webContext) Req() *http.Request {
	return this.Request
}

func (this *webContext) Param(name string) (string, bool) {
	value, ok := this.Params[name]
	return value, ok
}

// ------- Session Management -------

func (this *webContext) Session() framework.Session {
	if this.session == nil {
		panic("WebContext.Session cannot be nil")
	}

	return this.session
}

// ------- Result Management -------

func (this *webContext) Text(value string) framework.WebResult {

	return &TextResult{Value: value}
}

func (this *webContext) Binary(value []byte) framework.WebResult {

	return &StreamResult{
		Reader:  bytes.NewReader(value),
		Name:    "",
		ModTime: time.Now(),
		Length:  int64(len(value)),
		Inline:  true,
	}
}

func (this *webContext) Json(value interface{}) framework.WebResult {

	return &JsonResult{Value: value}
}

func (this *webContext) Xml(value interface{}) framework.WebResult {

	return &XmlResult{Value: value}
}

func (this *webContext) Error(err error) framework.WebResult {

	if fErr, ok := err.(*framework.FrameworkError); ok {
		return this.FrameworkError(fErr)
	}

	return &ErrorResult{Err: err}
}

func (this *webContext) FrameworkError(err *framework.FrameworkError) framework.WebResult {

	return &FrameworkErrorResult{Err: err}
}

func (this *webContext) Template(template *template.Template, data interface{}) framework.WebResult {

	return &TemplateResult{Template: template, Data: data}
}

func (this *webContext) NotFound(message string) framework.WebResult {

	return &GenericResult{Status: http.StatusNotFound, Message: message}
}

func (this *webContext) Forbidden(message string) framework.WebResult {
	return &GenericResult{Status: http.StatusForbidden, Message: message}
}

func (this *webContext) Unauthorized(message string) framework.WebResult {
	return &GenericResult{Status: http.StatusUnauthorized, Message: message}
}

func (this *webContext) Redirect(url string) framework.WebResult {
	return &RedirectResult{Url: url, Status: http.StatusFound}
}

func NewContext(server *webServer, wrapped *_web.Context) *webContext {

	context := &webContext{server: server}
	context.Context = wrapped

	return context
}
