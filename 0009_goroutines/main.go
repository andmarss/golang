package main

import (
	"fmt"
	"time"
)

/*
executing a function as a goroutine (place "go" before function name):

go funcName()

executing an anonymous function as a goroutine

go func()
{
 // function implementation goes here
}()
*/

func printGreetings(source string) {
	for i := 0; i < 9; i++ {
		fmt.Println(i, source)
	}
	time.Sleep(time.Millisecond * 1)
}

func main() {
	go printGreetings("goroutine")
	printGreetings("main function")
}
