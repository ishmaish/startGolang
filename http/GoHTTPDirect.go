package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirected to:", req.URL)
			return nil
		},
	}

	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
}
