package web

import (
	"einvite/framework"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"
)

//---------- ERROR -----------
type ErrorResult struct {
	Err error
}

func (r *ErrorResult) Write(ctx framework.WebContext) error {

	//TODO: implement properly
	ctx.Write([]byte("Here goes ErrorResult: " + r.Err.Error()))
	return nil
}

//---------- Framework Error -----------
type FrameworkErrorResult struct {
	Err *framework.FrameworkError
}

func (r *FrameworkErrorResult) Write(ctx framework.WebContext) error {
	//TODO: implement properly
	ctx.Write([]byte("Here goes FrameworkErrorResult: " + r.Err.Error()))
	return nil
}

//---------- JSON -----------
type JsonResult struct {
	Value interface{}
}

var jsonpEnclosing = []byte(");")
var jsonpFormat = "%s("

func (r *JsonResult) Write(ctx framework.WebContext) error {
	data, err := json.Marshal(r.Value)

	if err != nil {
		return err
	}

	ctx.WriteHeader(http.StatusOK)

	callback, _ := ctx.Param("callback")

	if callback == "" {
		//json
		ctx.ContentType("application/json")
		_, err = ctx.Write(data)
	} else {
		//jsonp
		ctx.ContentType("application/javascript")
		_, err = ctx.Write([]byte(fmt.Sprintf(jsonpFormat, callback)))
		_, err = ctx.Write(data)
		_, err = ctx.Write(jsonpEnclosing)
	}

	return err
}

//---------- XML -----------
type XmlResult struct {
	Value interface{}
}

func (r *XmlResult) Write(ctx framework.WebContext) error {
	data, err := xml.Marshal(r.Value)

	if err == nil {
		ctx.WriteHeader(http.StatusOK)
		ctx.ContentType("application/xml")
		ctx.Write(data)
	}

	return err
}

//---------- Text -----------
type TextResult struct {
	Value string
}

func (r *TextResult) Write(ctx framework.WebContext) error {

	ctx.WriteHeader(http.StatusOK)
	ctx.ContentType("text/plain")
	_, err := ctx.Write([]byte(r.Value))

	return err
}

type StreamResult struct {
	Reader  io.Reader
	Name    string
	Length  int64
	ModTime time.Time
	Inline  bool
}

//adapted from Revel framework BinaryResult.Apply method
func (r *StreamResult) Write(ctx framework.WebContext) error {

	var err error

	// If we have a ReadSeeker, delegate to http.ServeContent
	if rs, ok := r.Reader.(io.ReadSeeker); ok {
		http.ServeContent(ctx, ctx.Req(), r.Name, r.ModTime, rs)

	} else {
		// Else, do a simple io.Copy.

		if r.Length != -1 {
			ctx.Header().Set("Content-Length", strconv.FormatInt(r.Length, 10))
		}

		ctx.WriteHeader(http.StatusOK)
		//TODO: use a mime-type to file extension strategy like Revel does
		//ctx.ContentType("application/octet-stream")
		//ctx.ContentType(ContentTypeByFilename(r.Name))
		_, err = io.Copy(ctx, r.Reader)
	}

	// Close the Reader if we can
	if v, ok := r.Reader.(io.Closer); ok {
		v.Close()
	}

	return err
}

type TemplateResult struct {
	Template *template.Template
	Data     interface{}
}

func (r *TemplateResult) Write(ctx framework.WebContext) error {

	err := r.Template.Execute(ctx, r.Data)
	return err
}

type RedirectResult struct {
	Url    string
	Status int
}

func (r *RedirectResult) Write(ctx framework.WebContext) error {

	ctx.Header().Set("Location", r.Url)
	ctx.WriteHeader(r.Status)
	return nil
}

type GenericResult struct {
	Status  int
	Message string
}

func (r *GenericResult) Write(ctx framework.WebContext) error {

	if r.Status > 0 {
		ctx.WriteHeader(r.Status)
	} else {
		ctx.WriteHeader(http.StatusOK)
	}

	if r.Message != "" {
		ctx.Write([]byte(r.Message))
	}

	return nil
}
