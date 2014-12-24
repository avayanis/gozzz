package goz

import (
	"net/http"
	"strings"
)

// Request provides convenience functions on top of the standard http.Request
// interface.
type Request struct {
	host          string
	requestParams map[string]string

	*http.Request
}

// NewRequest constructs and returns an initialized Request
func NewRequest(request *http.Request) *Request {
	gozRequest := new(Request)

	gozRequest.Request = request

	// Extract host information from request
	gozRequest.host = strings.Split(request.Host, ":")[0]
	gozRequest.requestParams = make(map[string]string)

	return gozRequest
}

// SetVariableMap creates a variable map using the supplied keys and values.
func (request *Request) SetVariableMap(keys map[string]string, values map[string]string) {
	// Loop through all values and add them to the requestParams.  If a value
	// does not have a corresponding key, just use the index value for the key.
	for i := range values {
		if _, ok := keys[i]; ok {
			request.requestParams[keys[i][1:]] = values[i]
		} else {
			request.requestParams[i] = values[i]
		}
	}
}

// Host is an accessor method to Request.host
func (request *Request) Host() string {
	return request.host
}

// RequestParams is an accessor method to Request.host
func (request *Request) RequestParams() map[string]string {
	return request.requestParams
}
