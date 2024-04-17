package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a new HTTP client
	client := &http.Client{}

	// create a new GET request
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// check if the response was a redirect
	if resp.StatusCode >= 300 && resp.StatusCode <= 399 {
		redirectUrl, err := resp.Location()
		if err != nil {
			fmt.Println("Error getting redirect location:", err)
			return
		}

		// create a new GET request to follow the redirect
		req.URL = redirectUrl
		resp, err = client.Do(req)
		if err != nil {
			fmt.Println("Error sending redirect request:", err)
			return
		}
		defer resp.Body.Close()
	}

	// process the response
	fmt.Println("Status code:", resp.StatusCode)
	// ...
}
