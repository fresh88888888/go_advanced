package app

import (
	"log"
	"net/http"
)

type HandleFunc func(ctx *Context)

type App struct {
	router     *Router
	middleware []HandleFunc
	Context    Context
	config     *Config
}

func New() *App {
	a := new(App)
	a.router = NewRouter()
	a.config = LoadConfig()
	a.middleware = make([]HandleFunc, 0)

	return a
}

func (app *App) Use(h ...HandleFunc) {
	app.middleware = append(app.middleware, h...)
}

func (app *App) Config() *Config {
	return app.config
}

func (app *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	app.doHandle(res, req)
}

func (app *App) doHandle(res http.ResponseWriter, req *http.Request) {
	context := NewContext(res, req, app)

	if len(app.middleware) > 0 {
		for _, h := range app.middleware {
			h(context)
			if context.DidSent {
				break
			}
		}
	}

	if context.DidSent {
		return
	}
	router, err := app.router.MatchRoute(req.URL.Path, req.Method)
	if err != nil {
		context.SendStatus(http.StatusNotFound)
	}

	if router != nil {
		router.h(context)
		if context.DidSent {
			return
		}
	} else {
		context.SendStatus(http.StatusNotFound)
	}
}

func (app *App) Run() {
	addr := app.config.Addr
	log.Fatal(http.ListenAndServe(addr, app))
}

func (app *App) Post(pattern string, handler HandleFunc) {
	r := app.router.Route(pattern, http.MethodPost, handler)
	app.router.routes = append(app.router.routes, r)
}

func (app *App) Get(pattern string, handler HandleFunc) {
	r := app.router.Route(pattern, http.MethodGet, handler)
	app.router.routes = append(app.router.routes, r)
}

func (app *App) Put(pattern string, handler HandleFunc) {
	r := app.router.Route(pattern, http.MethodPut, handler)
	app.router.routes = append(app.router.routes, r)
}

func (app *App) Delete(pattern string, handler HandleFunc) {
	r := app.router.Route(pattern, http.MethodDelete, handler)
	app.router.routes = append(app.router.routes, r)
}

func (app *App) Patch(pattern string, handler HandleFunc) {
	r := app.router.Route(pattern, http.MethodPatch, handler)
	app.router.routes = append(app.router.routes, r)
}
