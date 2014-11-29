package gozzz

import (
	"net/http"
)

type GoAppHandlerFunc func(http.ResponseWriter, *http.Request) error

// NotFoundRoute returns a Route with a 404 handler.
func NotFoundRoute() *Route {
	route := new(Route)

	route.SetHandler(notFoundHandler)

	return route
}

func notFoundHandler(res http.ResponseWriter, req *http.Request) error {
	http.Error(res, http.StatusText(404), 404)

	return nil
}

func serverErrorHandler(res http.ResponseWriter, req *http.Request) error {
	http.Error(res, http.StatusText(500), 500)

	return nil
}
