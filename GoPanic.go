package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// In Go (Golang), a panic is a mechanism that allows you to halt the normal execution
// of a program when an unexpected or unrecoverable situation occurs.
// It is similar to an exception in other programming languages.

type Config struct {
	SecretKey string `json:"secret_key"`
}

func main() {
	// Read configuration file.
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		// If there's an error reading the file, panic and terminate the program.
		panic(fmt.Sprintf("Error reading configuration file: %v", err))
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		// If there's an error unmarshaling the JSON data, panic and terminate the program.
		panic(fmt.Sprintf("Error parsing configuration data: %v", err))
	}

	if config.SecretKey == "" {
		// If the secret key is empty, panic and terminate the program.
		panic("Secret key is missing from configuration")
	}

	// Continue with the program execution.
	fmt.Println("Application started successfully.")
}

// signature of panic function
//func panic(v interface{})
//The panic() function takes a single argument v of the empty interface type (interface{}).
//This means that you can pass any value of any type as an argument to panic().
//When a panic is triggered, the provided value is typically used to describe
//the reason for the panic,
//such as an error message or an error object.
