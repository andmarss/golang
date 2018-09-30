package main

import "fmt"

func main() {
	//buffered channel (with fixed size - 3)
	messageQueue := make(chan interface{}, 3)
	messageQueue <- "one"
	messageQueue <- 2
	messageQueue <- func() {}

	close(messageQueue)

	for m := range messageQueue {
		fmt.Println(m)
	}
}
