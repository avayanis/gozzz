package goz

// GoAppHandlerFunc extends http.Handle and adds a return parameter for errors.
type GoAppHandlerFunc func(ResponseWriter, *Request) error

// notFoundHandler responds with a standard 404 Not Found response.
func notFoundHandler(res ResponseWriter, req *Request) error {
	res.Error(404)

	return nil
}

// serverErrorHandler responds with a standard 500 Internal Server Error
// response.
func serverErrorHandler(res ResponseWriter, req *Request) error {
	res.Error(500)

	return nil
}
