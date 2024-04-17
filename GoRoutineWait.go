package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

/*
The WaitGroup type of sync package, is used to wait for the program to finish all goroutines
launched from the main function. It uses a counter that specifies the number of goroutines,
and Wait blocks the execution of the program until the WaitGroup counter is zero.
*/

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

/*The Add method is used to add a counter to the WaitGroup.

The Done method of WaitGroup is scheduled using a defer statement to decrement the WaitGroup counter.

The Wait method of the WaitGroup type waits for the program to finish all goroutines.*/

func main() {
	// Add a count of three, one for each goroutine.
	wg.Add(3)
	fmt.Println("Start Goroutines")

	go responseSize2("https://www.golangprograms.com")
	go responseSize2("https://stackoverflow.com")
	go responseSize2("https://coderwall.com")

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Terminating Program")
}

func responseSize2(url string) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}
