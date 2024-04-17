package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create an HTTP client
	client := &http.Client{}

	// Send an HTTP request
	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Read headers from the HTTP response
	contentType := resp.Header.Get("Content-Type")
	server := resp.Header.Get("Server")

	// Print headers
	fmt.Println("Content-Type:", contentType)
	fmt.Println("Server:", server)
}
