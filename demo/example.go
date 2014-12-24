package main

import (
	"github.com/avayanis/gozzz"
)

func main() {
	app := goz.NewGoApp(5050)

	// Set up route using AddRoute primitive
	app.AddRoute("GET", "/hello", func(response *goz.ResponseWriter, request *goz.Request) error {
		response.Write([]byte("world!"))

		return nil
	})

	// Set up route using convenience methods
	app.Get("/foo", func(response *goz.ResponseWriter, request *goz.Request) error {
		response.Write([]byte("bar!"))

		return nil
	})

	app.Post("/foo", func(response *goz.ResponseWriter, request *goz.Request) error {
		response.Write([]byte("You posted to me boo!"))

		return nil
	})

	// Set up dyamic route
	app.Get("/static/:dynamic", func(response *goz.ResponseWriter, request *goz.Request) error {
		response.Write([]byte("I am: " + request.RequestParams()["dynamic"]))

		return nil
	})

	app.Start()
}
