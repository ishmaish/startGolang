package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// recover() is a built-in function in Go that is used to regain control of a panicking goroutine.
// When a panic() is called, the normal flow of the program is interrupted,
// and the deferred functions in the same goroutine are executed.
// You can use recover() within a deferred function to catch the panic value,
// handle the error, and prevent the program from crashing.
func main() {

	// signature of recover func
	//func recover() interface{}

	/*The recover() function, when called inside a deferred function, returns the value that was passed to panic().
	If recover() is called outside a deferred function or if no panic() occurred, it returns nil.*/
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error occurred: %v\n", r)
			fmt.Println("Application terminated gracefully.")
		} else {
			fmt.Println("Application executed successfully.")
		}
	}()

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
