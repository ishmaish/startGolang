package main

import "fmt"

//In Go, it is possible to change the channel being selected on in a for-select loop.
//To do this, you can use a default case to select a new channel.

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", x)
			} else {
				fmt.Println("ch1 closed")
				ch1 = nil // set ch1 to nil to stop receiving from it
			}

		case x, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", x)
			} else {
				fmt.Println("ch2 closed")
				ch2 = nil // set ch2 to nil to stop receiving from it
			}

		default:
			// select a new channel to receive from
			if ch1 == nil && ch2 == nil {
				// both channels closed, exit the loop
				return
			} else if ch1 == nil {
				fmt.Println("Waiting for ch2")
				<-ch2
			} else if ch2 == nil {
				fmt.Println("Waiting for ch1")
				<-ch1
			} else {

				select {
				case x, ok := <-ch1:
					if ok {
						fmt.Println("Received from ch1:", x)
					} else {
						fmt.Println("ch1 closed")
						ch1 = nil
					}

				case x, ok := <-ch2:
					if ok {
						fmt.Println("Received from ch2:", x)
					} else {
						fmt.Println("ch2 closed")
						ch2 = nil
					}
				}
			}
		}
	}
}
