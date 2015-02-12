package goz

import (
	"reflect"
	"testing"
)

func TestNewRoute(t *testing.T) {
	expectedID := "/hello"
	expectedTypeOfChildren := "map[string]*goz.Route"
	expectedTypeOfHandlers := "map[string]goz.GoAppHandlerFunc"
	expectedTypeOfVariableMap := "map[string]map[string]string"

	route := NewRoute(expectedID)

	if expectedID != route.ID {
		t.Errorf("Expected route.ID to be %s, but got %s instead.", expectedID, route.ID)
	}

	typeOfChildren := reflect.TypeOf(route.children).String()
	if expectedTypeOfChildren != typeOfChildren {
		t.Errorf("Expected route.children to be %s, but got %s instead.", expectedTypeOfChildren, typeOfChildren)
	}

	typeOfHandlers := reflect.TypeOf(route.handlers).String()
	if expectedTypeOfHandlers != typeOfHandlers {
		t.Errorf("Expected route.children to be %s, but got %s instead.", expectedTypeOfHandlers, typeOfHandlers)
	}

	typeOfVariableMap := reflect.TypeOf(route.variableMap).String()
	if expectedTypeOfVariableMap != typeOfVariableMap {
		t.Errorf("Expected route.children to be %s, but got %s instead.", expectedTypeOfVariableMap, typeOfVariableMap)
	}
}

func TestRouteAddStaticRoute(t *testing.T) {
	route := NewRoute("")
	newRouteID := "test"
	newRoute := route.AddRoute(newRouteID)

	if newRouteID != newRoute.ID {
		t.Errorf("Expected new route.ID to be %s, but got %s instead.", newRouteID, newRoute.ID)
	}

	if route.children[newRouteID] != newRoute {
		t.Errorf("Expected route to have newRoute as child")
	}
}

func TestRouteAddDuplicateStaticRoute(t *testing.T) {
	route := NewRoute("")
	newRouteID := "test"
	route.AddRoute(newRouteID)

	var expectedNewRoute *Route
	newRoute := route.AddRoute(newRouteID)

	if newRoute != expectedNewRoute {
		t.Errorf("Adding duplicate route should return %s, but got %s instead.", expectedNewRoute, newRoute)
	}
}

func TestRouteAddDynamicRoute(t *testing.T) {
	route := NewRoute("")
	newRouteID := ":test"
	expectedRouteID := ":var"
	newRoute := route.AddRoute(newRouteID)

	if expectedRouteID != newRoute.ID {
		t.Errorf("Expected new route.ID to be %s, but got %s instead.", expectedRouteID, newRoute.ID)
	}

	if route.children[expectedRouteID] != newRoute {
		t.Errorf("Expected route to have newRoute as child")
	}
}

func TestRouteAddDuplicateDynamicRoute(t *testing.T) {
	route := NewRoute("")
	newRouteID := ":test"
	route.AddRoute(newRouteID)

	var expectedNewRoute *Route
	newRoute := route.AddRoute(newRouteID)

	if newRoute != expectedNewRoute {
		t.Errorf("Adding duplicate route should return %s, but got %s instead.", expectedNewRoute, newRoute)
	}
}

func TestRouteGetValidRoute(t *testing.T) {
	routeID := "/hello"
	expectedID := "world"
	route := NewRoute(routeID)
	expectedRoute := route.AddRoute(expectedID)

	if expectedRoute != route.GetRoute(expectedID) {
		t.Errorf("GetRoute did not return the correct route")
	}
}

func TestRouteGetInvalidRoute(t *testing.T) {
	routeID := "/hello"
	expectedID := "world"
	route := NewRoute(routeID)
	route.AddRoute(expectedID)

	var expectedRoute *Route
	if expectedRoute != route.GetRoute("shouldnotexist") {
		t.Errorf("GetRoute should return empty *Route when route does not exist")
	}
}

func TestRouteSethandler(t *testing.T) {
	routeID := "/hello"
	route := NewRoute(routeID)
	handlerMethod := "GET"

	var expectedHandler GoAppHandlerFunc

	expectedHandler = func(response ResponseWriter, request *Request) error {
		return nil
	}

	route.SetHandler(handlerMethod, expectedHandler)

	handler := route.Handler(handlerMethod)

	if reflect.ValueOf(expectedHandler) != reflect.ValueOf(handler) {
		t.Errorf("Handler Getter did not return what was set SetHandler")
	}
}

func TestRouteGetHandlerDoesNotExist(t *testing.T) {
	routeID := "/hello"
	route := NewRoute(routeID)
	handlerMethod := "GET"

	handler := route.Handler(handlerMethod)

	if nil != handler {
		t.Errorf("Route should return nil handler when handler does not exist.")
	}
}

func TestRouteSetVariableMap(t *testing.T) {
	routeID := "/hello"
	route := NewRoute(routeID)
	handlerMethod := "GET"
	expectedVariableMap := make(map[string]string)

	route.SetVariableMap(handlerMethod, expectedVariableMap)

	variableMap := route.VariableMap(handlerMethod)

	if reflect.ValueOf(expectedVariableMap) != reflect.ValueOf(variableMap) {
		t.Errorf("VariableMap Getter did not return what was set by SetVariableMap")
	}
}

func TestRouteGetVariableMapDoesNotExist(t *testing.T) {
	routeID := "/hello"
	route := NewRoute(routeID)
	handlerMethod := "GET"

	variableMap := route.VariableMap(handlerMethod)

	if nil != variableMap {
		t.Errorf("Route should return nil variableMap when variableMap does not exist.")
	}
}
