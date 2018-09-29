package main

import (
	"../0001_packages"
	"fmt"
)

/*
	Numeric Types
	Integers: int, int8, int16, int32, int64, uint, uint8, uin16, uint32, uint64
	Floating-Point numbers: float32, float64
	Complex Numbers: complex64, complex128

	uint8: 0 to 255
	uint16: 0 to 65535
	uint32: 0 to 4294967295
	uint64: 0 to 18446744073709551615

	int8: -128 to 127
	int16: -32768 to 32767
	int32: -2147483648 to 2147483647
	int64: -9223372036854775808 to ...
*/

const (
	LosAngeles = 1984 + (iota * 4)
	Seoul
	Barselona
	Atlanta
	Sydney
	Athens
	Beijing
	London
	Rio
	Tokyo
)

func main() {
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

	fmt.Println("These cities hosted or will host the Summer Olympics in the provider year...")

	fmt.Printf("%-18s %-18v \n", "Лос Анджелес", LosAngeles)
	fmt.Printf("%-18s %-18v \n", "Seoul", Seoul)
	fmt.Printf("%-18s %-18v \n", "Barselona", Barselona)
	fmt.Printf("%-18s %-18v \n", "Atlanta", Atlanta)
	fmt.Printf("%-18s %-18v \n", "Sydney", Sydney)
	fmt.Printf("%-18s %-18v \n", "Athens", Athens)
	fmt.Printf("%-18s %-18v \n", "Beijing", Beijing)
	fmt.Printf("%-18s %-18v \n", "London", London)
	fmt.Printf("%-18s %-18v \n", "Rio", Rio)
	fmt.Printf("%-18s %-18v \n", "Tokyo", Tokyo)

	greetingpackage.PrintGreetings()

	func() {
		fmt.Println("immediate anonymous call function")
	}()
}
