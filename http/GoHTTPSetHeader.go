package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP request with custom headers
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	req.Header.Add("Authorization", "Bearer <token>")
	req.Header.Add("Content-Type", "application/json")

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
