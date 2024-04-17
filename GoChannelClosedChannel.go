package main

import "fmt"

/*In Golang, if a channel has been closed, any further attempts to send values on that channel
will result in a panic. However, receiving from a closed channel is a safe operation,
and it is possible to receive values
from a closed channel until all the values in the channel have been read.*/

//When a channel is closed, it is guaranteed that all the values that were previously sent on that
//channel will be received,
//in the order they were sent, before the channel is closed.

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for {
		i, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Received", i)
	}
}

/*Note that it is important to use the ok variable to check if the channel is still open,
as attempting to receive from a closed channel without this check can result in a panic.*/
