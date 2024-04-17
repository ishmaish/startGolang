package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*To handle HTTP authentication with an HTTP client in Go,
you can set the Authorization header in the http.Request object.
There are several types of HTTP authentication, including Basic, Digest, and Bearer.*/

func main() {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "http://httpbin.org/basic-auth/user/passwd", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the HTTP Basic authentication header
	username := "user"
	password := "passwd"
	auth := username + ":" + password
	base64Auth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+base64Auth)

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP Error:", resp.Status)
		return
	}

	// Read the HTTP response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the HTTP response body
	fmt.Println("Response Body:", string(body))
}
