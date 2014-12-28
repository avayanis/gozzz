package goz

import (
	"net/http"
)

// ResponseWriter is an extension of the http.ResponseWriter interface that
// provides additional capabilities.
type ResponseWriter interface {
	http.ResponseWriter
	Error(int)
}

// Response is a goz extension of the http.Response
type Response struct {
	http.ResponseWriter
}

// NewResponse constructs and returns an initialized goz.Response
func NewResponse(writer http.ResponseWriter) *Response {
	response := &Response{writer}

	return response
}

// Error is an alias to http.Error
func (res *Response) Error(code int) {
	http.Error(res, http.StatusText(code), code)
}
