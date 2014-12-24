package goz

import (
	"net/http"
)

// ResponseWriter is a goz extension of the http.ResponseWriter
type ResponseWriter struct {
	writer http.ResponseWriter
}

// NewResponseWriter constructs and returns an initialized goz.Response
func NewResponseWriter(writer http.ResponseWriter) *ResponseWriter {
	response := new(ResponseWriter)

	response.writer = writer

	return response
}

// Header is an alias to writer.Header()
func (res *ResponseWriter) Header() map[string][]string {
	return res.writer.Header()
}

// WriteHeader is an alias to writer.Write([]byte)
func (res *ResponseWriter) Write(input []byte) (int, error) {
	return res.writer.Write(input)
}

// WriteHeader is an alias to writer.WriteHeader(int)
func (res *ResponseWriter) WriteHeader(statusCode int) {
	res.writer.WriteHeader(statusCode)
}
