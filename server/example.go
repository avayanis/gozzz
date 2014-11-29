package main

import (
	"github.com/avayanis/gozzz"
	"net/http"
)

func testHandler(response http.ResponseWriter, request *http.Request) error {
	response.Write([]byte("Here I am!"))

	return nil
}

func main() {
	app := gozzz.NewGoApp(5050)

	app.AddRoute("GET", "/test", testHandler)
	app.AddRoute("GET", "/test/:var", testHandler)
	app.AddRoute("GET", "/test/:var/test2", testHandler)

	app.Start()
}
