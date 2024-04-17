package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// send a GET request
	resp, err := http.Get("http://example.com")
	if err != nil {
		// handle error
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("Error reading response body:", err)
		return
	}

	// check for errors in the response
	if resp.StatusCode != http.StatusOK {
		// handle error
		fmt.Println("Error:", resp.StatusCode, string(body))
		return
	}

	// process the response
	// ...
}
