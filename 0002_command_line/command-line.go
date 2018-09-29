package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Andrew", "Programmer name")
	flag.Parse()
	fmt.Println("Hello " + name)
}
