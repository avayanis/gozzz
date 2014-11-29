package gozzz

import (
	"fmt"
	"net/http"
	"strings"
)

// Router stores a routing table and provides methods for adding routes and
// searching for routes.
type Router struct {
	routes map[string]*Route
}

// NewRouter returns an initialized Router
func NewRouter() *Router {
	router := new(Router)

	router.routes = make(map[string]*Route)

	return router
}

// ServeHTTP implements the http.Handler interface and acts as the generic
// entry to all incoming requests.
func (router Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var host string
	var route *Route

	hostInfo := strings.Split(req.Host, ":")
	pathSegments := parseRoute(req.URL.String())

	host = hostInfo[0]

	if _, ok := router.routes[host]; ok {
		route = routeRequest(router.routes[host], pathSegments[1:])
	} else if _, ok := router.routes["*"]; ok {
		route = routeRequest(router.routes["*"], pathSegments[1:])
	} else {
		route = new(Route)
	}

	handler := route.Handler()

	if handler != nil {
		err := handler(res, req)

		if err != nil {
			serverErrorHandler(res, req)
		}
	} else {
		notFoundHandler(res, req)
	}
}

// AddRoute adds a new route to the routing table
func (router *Router) AddRoute(method string, route string, handler GoAppHandlerFunc) {
	fmt.Println("Adding Route: " + route)
	pathSegments := parseRoute(route)

	segment := pathSegments[0]

	if segment == "" {
		segment = "*"
	}

	if _, ok := router.routes[segment]; !ok {
		router.routes[segment] = NewRoute(segment)
	}

	recurseInsertRoute(pathSegments[1:], router.routes[segment], handler)
}

func routeRequest(route *Route, pathSegments []string) *Route {
	if route.HasRoute(pathSegments[0]) {
		nextRoute := route.GetRoute(pathSegments[0])

		if len(pathSegments) == 1 {
			return nextRoute
		}

		return routeRequest(nextRoute, pathSegments[1:])
	}

	// Could not find a valid route.  Return an Empty Route
	return new(Route)
}

func recurseInsertRoute(pathSegments []string, route *Route, handler GoAppHandlerFunc) {
	var nextRoute *Route

	segment := pathSegments[0]
	if route.HasRoute(segment) {
		nextRoute = route.GetRoute(segment)
	} else {
		nextRoute = route.AddRoute(segment)
	}

	if len(pathSegments) > 1 {
		recurseInsertRoute(pathSegments[1:], nextRoute, handler)
	} else {
		nextRoute.SetHandler(handler)
	}
}

func parseRoute(route string) []string {
	return strings.Split(route, "/")
}
