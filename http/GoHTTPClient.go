package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// An HTTP client in Go is a program that sends HTTP requests to an HTTP server
// and receives HTTP responses
func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	// "defer" keyword to ensure that it is closed even if there is an error.
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
