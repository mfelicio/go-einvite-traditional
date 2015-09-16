package framework

import (
	"html/template"
	"net/http"
)

type WebContext interface {
	http.ResponseWriter
	Req() *http.Request

	//Always returns the Session
	//Should panic if nil
	//Its up to the WebContext/WebServer implementation to ensure that this method
	//can be safely invoked
	Session() Session

	Param(name string) (string, bool)
	ContentType(val string) string
	SetCookie(cookie *http.Cookie)

	Text(value string) WebResult
	Binary(value []byte) WebResult
	Json(value interface{}) WebResult
	Xml(value interface{}) WebResult
	Error(err error) WebResult
	FrameworkError(err *FrameworkError) WebResult
	Template(template *template.Template, data interface{}) WebResult
	//File(path string) WebResult

	NotFound(message string) WebResult
	Forbidden(message string) WebResult
	Unauthorized(message string) WebResult
	Redirect(url string) WebResult
}
