package app

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter

	// HTTP status code
	statusCode int

	// The application
	App *App

	// Did we already sent the resource
	DidSent bool
}

// NewContext creates a Golf.Context instance.
func NewContext(res http.ResponseWriter, req *http.Request, app *App) *Context {
	ctx := Context{
		Request:    req,
		Response:   res,
		statusCode: http.StatusOK,
		App:        app,
		DidSent:    false,
	}

	ctx.Request.ParseForm()
	return &ctx
}

func (ctx *Context) reset() {
	ctx.statusCode = http.StatusOK
	ctx.DidSent = false
}

func (ctx *Context) SendStatus(statusCode int) {
	ctx.statusCode = statusCode
	ctx.Response.WriteHeader(statusCode)
}

func (ctx *Context) StatusCode() int {
	return ctx.statusCode
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Response.Header().Set(key, value)
}

func (ctx *Context) AddHeader(key, value string) {
	ctx.Response.Header().Add(key, value)
}

func (ctx *Context) GetHeader(key string) string {
	return ctx.GetHeader(key)
}

func (ctx *Context) Query(key string, index ...int) (string, error) {
	if val, ok := ctx.Request.Form[key]; ok {
		if len(index) == 1 {
			return val[index[0]], nil
		}
		return val[0], nil
	}

	return "", errors.New("Query: key not found.")
}

// Json set json response with data and proper header
func (ctx *Context) Json(data interface{}) {
	json, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	ctx.SetHeader("Content-Type", "application/json")
	ctx.Send(json)
}

// Send the response immediately. Set `ctx.IsSet` to `true` to make sure that the response won't be sent twice.
func (ctx *Context) Send(body interface{}) {
	if ctx.DidSent {
		return
	}

	switch body.(type) {
	case []byte:
		ctx.Response.Write(body.([]byte))
	case string:
		ctx.Response.Write([]byte(body.(string)))
	default:
		log.Fatal("Send: Invalid body type")
	}
	ctx.DidSent = true
}

// Redirect method sets the response as a 302 redirection
func (ctx *Context) Redirect(url string) {
	ctx.SetHeader("Location", url)
	ctx.SendStatus(302)
}

// ContentType sets content-type string
func (ctx *Context) ContentType(contentType string) {
	ctx.ContentType(contentType)
}
