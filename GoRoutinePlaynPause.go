package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Using channels we can play and pause execution of goroutine.
	//A channel handles this communication by acting as a conduit between goroutines.

	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause"

	time.Sleep(1 * time.Second)
	command <- "Play"

	time.Sleep(1 * time.Second)
	command <- "Stop"

	wg.Wait()
}

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}
