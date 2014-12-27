package goz

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
func (router Router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	req := NewRequest(request)
	res := NewResponse(response)

	// Get host information from request
	pathSegments := parseRoute(req.URL.String())

	// Find the correct routing table using the host string, and route the
	// incoming request.
	var route *Route
	var variableMap map[string]string
	if _, ok := router.routes[req.host]; ok {
		route, variableMap = routeRequest(NewPacket(router.routes[req.host], pathSegments[1:]))
	} else if _, ok := router.routes["*"]; ok {
		route, variableMap = routeRequest(NewPacket(router.routes["*"], pathSegments[1:]))
	} else {
		// Couldn't find a routing table for the requested host, so we create an
		// empty route.
		route = new(Route)
		variableMap = nil
	}

	req.SetVariableMap(route.VariableMap(req.Method), variableMap)
	handler := route.Handler(req.Method)

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
func routeRequest(packet *Packet) (*Route, map[string]string) {
	var nextRoute *Route

	// Search for a static route, if we can't find one, search for a dynamic
	// route.
	nextRoute = packet.route.GetRoute(packet.segments[0])
	if nextRoute == nil {
		nextRoute = packet.route.GetRoute(":var")

		if nextRoute != nil {
			// Let's store the variable segment
			packet.AddVariable()
		}
	}

	if nextRoute != nil {
		// Route matched, so let's return it!
		if len(packet.segments) == 1 {
			return nextRoute, packet.variableMap
		}

		// Let's keep searching.
		nextPacket := incrementPacket(nextRoute, packet)
		return routeRequest(nextPacket)
	}

	// We couldn't find a valid route.  Return an Empty Route.
	return new(Route), nil
}

// incremenetPacket creates a new goz.Packet that represents the next pointer
// in the routing table tree.
func incrementPacket(nextRoute *Route, previousPacket *Packet) *Packet {
	packet := NewPacket(nextRoute, previousPacket.segments[1:])

	packet.index = previousPacket.index + 1
	packet.variableMap = previousPacket.variableMap

	return packet
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

	packet := NewPacket(router.routes[segment], pathSegments[1:])
	recurseAddRoute(strings.ToUpper(method), packet, handler)
}

// recurseAddRoute adds a route to the routing table by recursively traversing
// the routing table and extending the table.
func recurseAddRoute(method string, packet *Packet, handler GoAppHandlerFunc) {
	var nextRoute *Route

	segment := packet.segments[0]

	if strings.HasPrefix(segment, ":") {
		segment = ":var"
		packet.AddVariable()
	}

	nextRoute = packet.route.GetRoute(segment)

	if nextRoute == nil {
		nextRoute = packet.route.AddRoute(segment)
	}

	if len(packet.segments) > 1 {
		// Prepare next routing packet
		nextPacket := NewPacket(nextRoute, packet.segments[1:])
		nextPacket.index = packet.index + 1
		nextPacket.variableMap = packet.variableMap

		// Continue recursion
		recurseAddRoute(method, nextPacket, handler)
	} else {
		// We have reached the end of the route, so we set the request handler.
		nextRoute.SetHandler(method, handler)
		nextRoute.SetVariableMap(method, packet.variableMap)
	}
}

// parseRoute returns an array of routing segments.
func parseRoute(route string) []string {
	return strings.Split(route, "/")
}
