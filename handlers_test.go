package goz

import (
	"net/http"
	"testing"
)

type ResponseMock struct {
	ResponseWriter
	ErrorMock func(int)
}

func NewResponseMock() *ResponseMock {
	responseMock := new(ResponseMock)

	return responseMock
}

func (mock *ResponseMock) Error(code int) {
	mock.ErrorMock(code)
}

func TestNotFoundHandler(t *testing.T) {
	responseMock := NewResponseMock()
	expectedCode := 404

	responseMock.ErrorMock = func(code int) {
		if code != expectedCode {
			t.Errorf("code expected to be %d, got %d", expectedCode, code)
		}
	}

	notFoundHandler(responseMock, NewRequest(new(http.Request)))
}

func TestServerErrorHandler(t *testing.T) {
	responseMock := NewResponseMock()
	expectedCode := 500

	responseMock.ErrorMock = func(code int) {
		if code != expectedCode {
			t.Errorf("code expected to be %d, got %d", expectedCode, code)
		}
	}

	serverErrorHandler(responseMock, NewRequest(new(http.Request)))
}
