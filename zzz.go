package goz

import (
	"net/http"
	"strconv"
)

// GoApp defines an HTTP Server with a RESTful interface.
type GoApp struct {
	port       int
	hostname   string
	router     *Router
	httpServer http.Server
}

// NewGoApp constructs and returns an initialized GoApp.
func NewGoApp(port int) *GoApp {
	app := new(GoApp)

	app.port = port
	app.router = NewRouter()
	app.httpServer.Addr = app.hostname + ":" + strconv.Itoa(app.port)
	app.httpServer.Handler = app.router

	return app
}

// Start will start up the built in HTTP Server.
func (app GoApp) Start() {
	app.httpServer.ListenAndServe()
}

// Get is a convenience function for adding HTTP GET routes
func (app GoApp) Get(route string, handler GoAppHandlerFunc) {
	app.router.AddRoute("GET", route, handler)
}

// Post is a convenience function for adding HTTP POST routes
func (app GoApp) Post(route string, handler GoAppHandlerFunc) {
	app.router.AddRoute("POST", route, handler)
}

// Update is a convenience function for adding HTTP UPDATE routes
func (app GoApp) Update(route string, handler GoAppHandlerFunc) {
	app.router.AddRoute("UPDATE", route, handler)
}

// Delete is a convenience function for adding HTTP DELETE routes
func (app GoApp) Delete(route string, handler GoAppHandlerFunc) {
	app.router.AddRoute("DELETE", route, handler)
}

// AddRoute registers a request handler
func (app GoApp) AddRoute(method string, route string, handler GoAppHandlerFunc) {
	app.router.AddRoute(method, route, handler)
}
