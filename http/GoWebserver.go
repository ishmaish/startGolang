package main

import (
	"fmt"
	"net/http"
)

/*A web server is a computer system, software, or a combination of both,
that hosts and delivers web content (such as web pages, images, videos,
and other resources) over the internet using the Hypertext Transfer Protocol (HTTP) or
its secure version, HTTPS. Web servers process incoming requests from clients (typically web browsers)
and respond with the requested content or an appropriate status message.*/

func main() {
	// PRIMARY FUNCTIONS OF A WEBSERVER
	//1. Listening for incoming req from clients ex: browser
	//2. Processing the req ex: image/webpage
	//3. sending a res to the req in the form of HTTP response ex: header(meta) and body

	//The web server is started by calling the main() function
	http.HandleFunc("/", hello)

	// the server registers a function called hello as the handler for the root path ("/")
	//using http.HandleFunc("/", hello). This means that when an HTTP request is received with
	//the root path,
	//the hello function will be called to handle it.

	http.ListenAndServe("0.0.0.0:8080", nil)

	//server starts listening for incoming HTTP requests on all available network interfaces ("0.0.0.0")
	//and port 8080 using http.ListenAndServe("0.0.0.0:8080", nil).
}

// the server calls the hello function to handle the request.
func hello(w http.ResponseWriter, _ *http.Request) {
	// the server writes "Hello" to the response using fmt.Fprintf(w, "Hello").
	//The http.ResponseWriter (w) is an interface that allows the server to construct an HTTP response,
	//and fmt.Fprintf writes the formatted string "Hello" to it.
	fmt.Fprintf(w, "Hello")
}
