package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Create a new user object
	user := User{Name: "John Doe", Email: "johndoe@example.com"}

	// Marshal the user object to JSON
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response
	w.Write(data)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
