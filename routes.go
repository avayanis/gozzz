package gozzz

// Route defines a child node in a routing table tree.
type Route struct {
	ID       string
	children map[string]*Route
	handler  GoAppHandlerFunc
}

// NewRoute returns an initialized Route
func NewRoute(id string) *Route {
	route := new(Route)

	route.ID = id
	route.children = make(map[string]*Route)

	return route
}

// AddRoute inserts a child route
func (route *Route) AddRoute(segment string) *Route {
	if _, ok := route.children[segment]; ok {
		return nil
	}

	route.children[segment] = NewRoute(segment)

	return route.children[segment]
}

// GetRoute retrieves child route
func (route *Route) GetRoute(segment string) *Route {
	if _, ok := route.children[segment]; ok {
		return route.children[segment]
	}

	return nil
}

// HasRoute tests if child route exists
func (route *Route) HasRoute(segment string) bool {
	if _, ok := route.children[segment]; ok {
		return true
	}

	return false
}

// SetHandler is a setter for route.handler
func (route *Route) SetHandler(handler GoAppHandlerFunc) {
	route.handler = handler
}

// Handler is a getter for route.handler
func (route *Route) Handler() GoAppHandlerFunc {
	return route.handler
}
