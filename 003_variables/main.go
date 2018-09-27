package main

import "fmt"

func main(){
	var a int = 5
	var b, c string = "b", "c"
	var d bool = true
	var e string
	var f = 1

	fmt.Printf("%T - %v\n", a, a)
	fmt.Printf("%T - %v\n", b, b)
	fmt.Printf("%T - %v\n", c, c)
	fmt.Printf("%T - %v\n", d, d)
	fmt.Printf("%T - %v\n", e, e)
	fmt.Printf("%T - %v\n", f, f)

	g := 100
	h := true

	fmt.Printf("%T - %v\n", g, g)
	fmt.Printf("%T - %v\n", h, h)

	var zeroInt int
	var emptyString string
	var boolFalse bool

	fmt.Printf("%T - %v\n", zeroInt, zeroInt)
	fmt.Printf("%T - %v\n", emptyString, emptyString)
	fmt.Printf("%T - %v\n", boolFalse, boolFalse)
}
