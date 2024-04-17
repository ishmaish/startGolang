package main

import (
	"net/http"
	"net/http/cookiejar"
)

func main() {
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: "12345",
	}

	client := &http.Client{
		Jar:       &cookiejar.Jar{},
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest("GET", "https://www.example.com", nil)
	if err != nil {
		// handle error
	}

	client.Jar.SetCookies(req.URL, []*http.Cookie{cookie})

	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	// read response
}
