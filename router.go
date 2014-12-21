package gozzz

import (
	"net/http"
	"strings"
)

// Router stores a routing table and provides methods for adding routes and
// searching for routes.
type Router struct {
	routes map[string]*Route
}

// NewRouter constructs and returns an initialized Router.
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

	// Extract host information from request
	hostInfo := strings.Split(req.Host, ":")
	pathSegments := parseRoute(req.URL.String())
	host = hostInfo[0]

	// Find the correct routing table using the host string, and route the
	// incoming request.
	if _, ok := router.routes[host]; ok {
		route = routeRequest(router.routes[host], pathSegments[1:])
	} else if _, ok := router.routes["*"]; ok {
		route = routeRequest(router.routes["*"], pathSegments[1:])
	} else {
		// Couldn't find a routing table for the requested host, so we create an
		// empty route.
		route = new(Route)
	}

	handler := route.Handler()

	if handler != nil {
		// Pass request to the appropriate handler.
		err := handler(res, req)

		if err != nil {
			// We encourntered an error, so serve an 500 'Internal Server Error'.
			serverErrorHandler(res, req)
		}
	} else {
		// We didn't find an appropriate request handler, so serve a 404
		// 'Page Not Found'.
		notFoundHandler(res, req)
	}
}

// routeRequest performs a lookup for a request and returns an appropriate
// route.
func routeRequest(route *Route, pathSegments []string) *Route {
	var nextRoute *Route
	// Search for a static route, if we can't find one, search for a dynamic
	// route.
	nextRoute = route.GetRoute(pathSegments[0])
	if nextRoute == nil {
		nextRoute = route.GetRoute(":var")
	}

	if nextRoute != nil {

		if len(pathSegments) == 1 {
			return nextRoute
		}

		return routeRequest(nextRoute, pathSegments[1:])
	}

	// Could not find a valid route.  Return an Empty Route.
	return new(Route)
}

// AddRoute adds a new route to the routing table.
func (router *Router) AddRoute(method string, route string, handler GoAppHandlerFunc) {
	pathSegments := parseRoute(route)

	segment := pathSegments[0]

	// If not host is provided, we treat the route as a wildcard route.
	if segment == "" {
		segment = "*"
	}

	if _, ok := router.routes[segment]; !ok {
		router.routes[segment] = NewRoute(segment)
	}

	recurseAddRoute(pathSegments[1:], router.routes[segment], handler)
}

// recurseAddRoute adds a route to the routing table by recursively traversing
// the routing table and extending the table.
func recurseAddRoute(pathSegments []string, route *Route, handler GoAppHandlerFunc) {
	var nextRoute *Route

	segment := pathSegments[0]
	nextRoute = route.GetRoute(segment)

	if nextRoute == nil {
		nextRoute = route.AddRoute(segment)
	}

	if len(pathSegments) > 1 {
		// Continue recursion
		recurseAddRoute(pathSegments[1:], nextRoute, handler)
	} else {
		// We have reached the end of the route, so we set the request handler.
		nextRoute.SetHandler(handler)
	}
}

// parseRoute returns an array of routing segments.
func parseRoute(route string) []string {
	return strings.Split(route, "/")
}
