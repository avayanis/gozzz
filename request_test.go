package goz

import (
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	expectedHost := "test.com"
	expectedPort := "8080"

	mockRequest := &http.Request{Host: expectedHost + ":" + expectedPort}
	request := NewRequest(mockRequest)

	if request.Host() != expectedHost {
		t.Errorf("Expected request.host to be %s, but got %s", expectedHost, request.host)
	}
}

func TestSetVariableMap(t *testing.T) {
	keys := make(map[string]string)
	keys["0"] = ":testKey0"
	keys["2"] = ":testKey2"

	values := make(map[string]string)
	values["0"] = "testValue0"
	values["2"] = "testValue2"
	values["3"] = "testValue3"

	expectedKey0 := keys["0"][1:]
	expectedKey3 := "3"

	mockRequest := new(http.Request)
	request := NewRequest(mockRequest)

	request.SetVariableMap(keys, values)

	if request.RequestParams()[expectedKey0] != values["0"] {
		t.Errorf("Expecting request.requestParams['%s'] to be %s, but got %s",
			expectedKey0, values["0"], request.RequestParams()[expectedKey0])
	}

	if request.RequestParams()[expectedKey3] != values["3"] {
		t.Errorf("Expecting request.requestParams['%s'] to be %s, but got %s",
			expectedKey3, values["3"], request.RequestParams()[expectedKey3])
	}
}
