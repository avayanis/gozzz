package gozzz

import (
	"strings"
)

// Route defines a child node in a routing table tree.
type Route struct {
	ID       string
	children map[string]*Route
	handlers map[string]GoAppHandlerFunc
}

// NewRoute constructs and returns an initialized Route.
func NewRoute(id string) *Route {
	route := new(Route)

	route.ID = id
	route.children = make(map[string]*Route)
	route.handlers = make(map[string]GoAppHandlerFunc)

	return route
}

// AddRoute creates a new route and adds a child node if the provided segment
// has not already been added.
func (route *Route) AddRoute(segment string) *Route {
	nodeName := segment

	if strings.HasPrefix(nodeName, ":") {
		nodeName = ":var"
	}

	if _, ok := route.children[nodeName]; ok {
		return nil
	}

	route.children[nodeName] = NewRoute(nodeName)

	return route.children[nodeName]
}

// GetRoute searches child nodes for the provided segment and returns the
// associated route or nil.
func (route *Route) GetRoute(segment string) *Route {
	if _, ok := route.children[segment]; ok {
		return route.children[segment]
	}

	return nil
}

// SetHandler is a setter for route.handler.
func (route *Route) SetHandler(method string, handler GoAppHandlerFunc) {
	route.handlers[method] = handler
}

// Handler is a getter for route.handler.
func (route *Route) Handler(method string) GoAppHandlerFunc {
	if _, ok := route.handlers[method]; ok {
		return route.handlers[method]
	}

	return nil
}
