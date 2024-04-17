package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	/*A goroutine is a lightweight thread of execution in the Go programming language.
	It is similar to a thread in other programming languages,
	but it is managed by the Go runtime rather than the operating system.
	Goroutines allow concurrent execution of functions in a program, and they
	are designed to be efficient and scalable.*/

	//This makes it easy to write concurrent programs in Go that
	//take advantage of multiple CPU cores and can perform many tasks simultaneously.

	/*
		goroutines are managed by the Go runtime, they are automatically scheduled and can communicate with each other using channels.
		This makes it easy to write complex concurrent programs without worrying about low-level details such as locking and synchronization.*/

	//Channels provide a way for goroutines to send and receive values to and from each other,
	//and they are used to synchronize access to shared data.

	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	go responseSize("https://stackoverflow.com")
	time.Sleep(10 * time.Second)
	//time.Sleep in the main function which prevents the main goroutine from exiting before the responseSize goroutines can finish.

}
func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}
