package main

import (
	"fmt"
	"sync"
)

// A mutex is used to create a critical section around code
// that ensures only one goroutine at a time can execute that code section.
var (
	//counter int32          // counter is a variable incremented by all goroutines.
	//wg      sync.WaitGroup // wg is used to wait for the program to finish.
	mutex sync.Mutex // mutex is used to define a critical section of code.
)

func main() {
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment2("Python")
	go increment2("Go Programming Language")
	go increment2("Java")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)
}

func increment2(lang string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for i := 0; i < 3; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
	//A critical section defined by the calls to Lock() and Unlock()
	//protects the actions against the counter variable and reading the text of name variable.
}
