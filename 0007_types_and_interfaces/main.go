package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1. Declaring a type
type type struct {
	field1 type
	field2 type
}

2. Creating an instance of a type (order matters)
t := type{field1Value, field2Value}

3. Creating an instance of a type using field names
t := type{field1: field1Value, field2: field2Value}

4. We declare methods using a pointer receiver
func (r *type) methodName() returnType{

}

5. Declaring an interface
type interfaceName interface{
	methodName1()returnType1
	methodName2()returnType2
}

Note: A type implements an interface when it has defined all the method in the method list of an interface

We can use pieces of an implementation by embedding types, either within a struct or an interface

type Reader interface {
	Read(p []byte)(n int, err error)
}

type Writer interface {
	Write(p []byte)(n int, err error)
}

type ReaderWriter interface {
	Reader
	Writer
}

Every type in Go implements the empty interface

interface{}

Some Use Cases

1. A function that return a value of interface{} can return any type

2. We can store heterogeneous values in an array, slice, or map using interface{} type
*/

type Triangle struct {
	width  float64
	height float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Shape interface {
	Area() float64
}

func newTriangle(w float64, h float64) *Triangle {
	return &Triangle{width: w, height: h}
}

func newRectangle(w float64, h float64) *Rectangle {
	return &Rectangle{width: w, height: h}
}

func (t *Triangle) Area() float64 {
	return (t.height * t.width) / 2
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func ShapeArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rectangle := newRectangle(10, 6)
	triangle := newTriangle(3, 7)

	fmt.Println("Area of triangle:", ShapeArea(triangle))
	fmt.Println("Area of rectangle:", ShapeArea(rectangle))

	var slice = []interface{}{1, "2", []interface{}{1, 2, func() {}}, func() {}}

	fmt.Println(test(rectangle))
	fmt.Println(test(triangle))
	fmt.Println(test(slice))
	// 1 or 0
	fmt.Println(test(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(2)))
}

func test(whatever interface{}) interface{} {
	return whatever
}
