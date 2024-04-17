package main

import "fmt"

/*In Go, a channel is a built-in data structure that is used for communication and
synchronization between goroutines (concurrent functions). Channels are a fundamental feature
of the language that enable safe and efficient communication
and synchronization between goroutines (concurrently executing functions).*/

//Channels are created using the built-in make function and can be buffered or unbuffered.

/*"Unbuffered channels" -  block the sending goroutine until there is a corresponding receiver ready to receive
the value being sent. This means that data is guaranteed to be received
in the order it was sent, and that synchronization is built into the channel*/

/*"Buffered channels", on the other hand, can hold a limited number of values (determined by the buffer size),
and will only block the sending goroutine when the buffer is full.
This can allow for some additional concurrency,
but requires careful consideration to avoid deadlocks and other synchronization issues.*/

func main() {
	// creating a simple channel
	ch := make(chan int)

	//sending and receiving values from a simple channel
	ch <- 42  // send 42 into the channel
	x := <-ch // receive a value from the channel and assign it to x
	fmt.Println(x)

	// BUFFERED CHANNELS

	/*Buffered channels can be used to build many types of concurrent systems,
	such as pipelines and worker pools, where multiple producers and consumers operate at different rates.
	However,
	it's important to choose an appropriate buffer size to avoid deadlock or performance issues.*/

	ch2 := make(chan int, 10) // create a buffered channel with a capacity of 10
	go fib(10, ch2)           // generate the first 10 Fibonacci numbers in a separate goroutine

	// read values from the channel until it's closed
	for x := range ch {
		fmt.Println(x)
	}

}

func fib(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
