package main

import "fmt"

/*In Golang, channels can be used as function parameters to allow for communication between goroutines.
This is useful for passing data between goroutines,
as well as for signaling when a goroutine has finished executing.*/

func main() {

	ch := make(chan int)
	go sendData(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
func sendData(ch chan<- int) {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
}
