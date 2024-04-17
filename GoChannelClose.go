package main

import "fmt"

//1. Only the sender should close the channel:
/*It is important to remember that only the sender should close the channel.
Closing a channel indicates that no more values will be sent on the channel,
and any attempts to send on the channel will result in a panic.*/

//2. Use the range loop to receive values from the channel:
/*The range loop can be used to receive values from the channel until the channel is closed.
The loop will automatically terminate when the channel is closed,
and any values sent before the channel was closed will be received.*/

//3. Check if the channel is closed before receiving a value:
/*It is possible to check if the channel is closed before receiving a value by using a comma ok idiom.
The idiom returns two values, the value received from the channel and a boolean value that is true
if the channel is open or false if the channel is closed.*/

//4. Use a select statement to receive values from multiple channels:
/*If you are receiving values from multiple channels, you can use a select statement to receive values
until all channels are closed.
The select statement will block until a value is received from one of the channels or all channels are closed.*/

func main() {
	input := make(chan int)
	output := make(chan int)
	done := make(chan bool)

	go worker(input, output, done)

	// Send some values on the input channel
	for i := 0; i < 10; i++ {
		input <- i
	}

	// Close the input channel to signal the end of input
	close(input)

	// Receive values from the output channel
	for n := range output {
		fmt.Println(n)
	}

	// Signal the worker to exit
	done <- true
}

func worker(input chan int, output chan int, done chan bool) {
	for {
		select {
		case n := <-input:
			// Do some work on n
			output <- n * 2
		case <-done:
			close(output)
			return
		}
	}
}

/*In this example, the worker function receives values from the input channel,
performs some work, and sends the result on the output channel.
The done channel is used to signal the worker to exit when all input has been processed.
The main function sends some values on the input channel, closes the input channel,
and then receives values from the output channel using a range loop.
Finally, the done channel is used to signal the worker to exit.*/
