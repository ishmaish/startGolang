package main

import (
	"fmt"
	"net/http"
)

/*The default HTTP server in Go is the "net/http" package, which provides
an implementation of HTTP server
and client functionality. This package allows you to create a server that can listen for
incoming HTTP requests, handle those requests, and return an appropriate HTTP response.
The "net/http" package includes features such as routing, middleware,
and support for various HTTP methods and headers.
It is a powerful and flexible tool for building web applications in Go.*/

/*An HTTP server in Go is a program that listens for incoming HTTP requests and sends back HTTP responses.
In Go, the standard library provides a package called "net/http"
that allows developers to easily create HTTP servers.*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
