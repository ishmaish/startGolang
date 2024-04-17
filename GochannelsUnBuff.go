package main

import "fmt"

/*In Go, an unbuffered channel is a type of channel that has no buffer
and is used for synchronous communication between goroutines.
When a value is sent on an unbuffered channel,
the sending goroutine blocks until another goroutine receives the value from the channel.
Likewise, when a goroutine attempts to receive a value from an unbuffered channel,
it blocks until another goroutine sends a value to the channel.*/

//Unbuffered channels are useful for ensuring that communication between goroutines is synchronized and
//that data is transferred reliably.
//They can be used to build many types of concurrent systems, such as message passing systems and pipelines.

//Note that the <- operator is used to both send and receive values on the channel.
//When used on the "left-hand" side of an assignment, the <- operator "sends a value" on the channel.
//When used on the "right-hand" side of an assignment, it "receives a value" from the channel.

func main() {
	ch := make(chan int) // create an unbuffered channel
	go fib2(10, ch)      // generate the first 10 Fibonacci numbers in a separate goroutine

	// read values from the channel until it's closed
	for i := 0; i < 10; i++ {
		x := <-ch
		fmt.Println(x)
	}
}
func fib2(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
}
