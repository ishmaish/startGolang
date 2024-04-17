package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
The most natural way to fetch a value from a goroutine is channels.
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive
those values into another goroutine or in a synchronous function.
*/
// WaitGroup is used to wait for the program to finish goroutines.
//var wg sync.WaitGroup (commented for ref the same varible is available in other goroutine file)

func main() {
	nums := make(chan int) // Declare a unbuffered channel
	wg.Add(1)
	go responseSize3("https://www.golangprograms.com", nums)
	fmt.Println(<-nums) // Read the value from unbuffered channel
	wg.Wait()
	close(nums) // Closes the channel
}

func responseSize3(url string, nums chan int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Send value to the unbuffered channel
	nums <- len(body)
}
