package goz

import (
	"net/http"
)

// GoAppHandlerFunc extends http.Handle and adds a return parameter for errors.
type GoAppHandlerFunc func(*ResponseWriter, *Request) error

// notFoundHandler responds with a standard 404 Not Found response.
func notFoundHandler(res *ResponseWriter, req *Request) error {
	http.Error(res.writer, http.StatusText(404), 404)

	return nil
}

// serverErrorHandler responds with a standard 500 Internal Server Error
// response.
func serverErrorHandler(res *ResponseWriter, req *Request) error {
	http.Error(res.writer, http.StatusText(500), 500)

	return nil
}
