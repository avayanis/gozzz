package main

import (
	"github.com/avayanis/gozzz"
	"net/http"
)

func main() {
	app := gozzz.NewGoApp(5050)

	// Set up route using AddRoute primitive
	app.AddRoute("GET", "/hello", func(response http.ResponseWriter, request *http.Request) error {
		response.Write([]byte("world!"))

		return nil
	})

	// Set up route using convenience methods
	app.Get("/foo", func(response http.ResponseWriter, request *http.Request) error {
		response.Write([]byte("boo!"))

		return nil
	})

	app.Post("/foo", func(response http.ResponseWriter, request *http.Request) error {
		response.Write([]byte("You posted to me boo!"))

		return nil
	})

	// Set up dyamic route
	app.Get("/static/:dynamic", func(response http.ResponseWriter, request *http.Request) error {
		response.Write([]byte("I am dynamic!"))

		return nil
	})

	app.Start()
}
