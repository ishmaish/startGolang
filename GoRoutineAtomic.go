package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

//Race conditions occur due to unsynchronized access to shared
//resource and attempt to read and write to that resource at the same time.

// Atomic functions provide low-level locking mechanisms for synchronizing access to integers and pointers

// Atomic functions generally used to fix the race condition.
var (
	counter int32 // counter is a variable incremented by all goroutines.
	//wg      sync.WaitGroup // wg is used to wait for the program to finish.
)

func main() {
	// The functions in the atomic under sync packages provides support
	//to synchronize goroutines by "locking access to shared resources".
	wg.Add(3) // Add a count of two, one for each goroutine.

	go increment("Python")
	go increment("Java")
	go increment("Golang")

	wg.Wait() // Wait for the goroutines to finish.
	fmt.Println("Counter:", counter)
}
func increment(name string) {
	defer wg.Done() // Schedule the call to Done to tell main we are done.

	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched() // Yield the thread and be placed back in queue.
	}
}

/*The AddInt32 function from the atomic package synchronizes the adding of integer values by enforcing
that only one goroutine can perform and complete this add operation at a time*/
