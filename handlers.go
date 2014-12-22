package gozzz

import (
	"net/http"
)

// GoAppHandlerFunc extends http.Handle and adds a return parameter for errors.
type GoAppHandlerFunc func(http.ResponseWriter, *http.Request) error

// notFoundHandler responds with a standard 404 Not Found response.
func notFoundHandler(res http.ResponseWriter, req *http.Request) error {
	http.Error(res, http.StatusText(404), 404)

	return nil
}

// serverErrorHandler responds with a standard 500 Internal Server Error
// response.
func serverErrorHandler(res http.ResponseWriter, req *http.Request) error {
	http.Error(res, http.StatusText(500), 500)

	return nil
}
