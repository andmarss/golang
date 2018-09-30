package main

import (
	"fmt"
)

/*
channels are conduits (pipes) that you can use to pass values of a particular type from one goroutine to another

channels allow goroutines to share memory by communicating

we use the channel operator, <-, to send and receive values
Note: the data flows in the direction of the arrow

we can create a channel using built-in make function

channel = make(chan type)

buffered channels are asynchronous, sending and receiving messages through the channel will not block unless the channel is full

we can create a buffered channel using the built-in make function

ch := make(chan type, capacity)
*/

func writeMessageToChannel(message chan string) {
	message <- "Hello!"
}

var done chan bool = make(chan bool)

func printGreetings(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println(i, source)
	}

	if source == "goroutine" {
		done <- true
	}
}

func main() {
	fmt.Println("Channel demo")

	message := make(chan string)
	go writeMessageToChannel(message)

	fmt.Println("Greeting from the message channel:", <-message)

	close(message)

	go printGreetings("goroutine")
	printGreetings("main function")

	<-done

	messageQueue := make(chan interface{}, 3)
	messageQueue <- "one"
	messageQueue <- 2
	messageQueue <- func() {}

	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
}
