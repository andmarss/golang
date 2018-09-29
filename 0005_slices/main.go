package main

import "fmt"

/*
1. Declaring and initializing a slice
var slice []type = make([]type, length)
var slice []type

2. Setting a value in the slice
slice[index] = value
slice = append(slice, value1, value2, ... valueN)

3. Getting a value in the slice
slice[index]
*/

func main() {
	var slice []int
	fmt.Printf("Contents of slice: %v\n", slice)
	fmt.Printf("Length of slice: %v\n", len(slice))
	fmt.Printf("Capacity of slice: %v\n", cap(slice))

	slice = append(slice, 1, 2, 3, 4)

	fmt.Printf("Contents of slice: %v\n", slice)
	fmt.Printf("Length of slice: %v\n", len(slice))
	fmt.Printf("Capacity of slice: %v\n", cap(slice))

	slice = append(slice, 5)

	fmt.Printf("Contents of slice: %v\n", slice)
	fmt.Printf("Length of slice: %v\n", len(slice))
	fmt.Printf("Capacity of slice: %v\n", cap(slice))

	fmt.Printf("Slice from first element: %v\n", slice[1:])
	fmt.Printf("Slice from first to last element: %v\n", slice[1:len(slice)-1])
	fmt.Printf("Slice from first to second element: %v\n", slice[1:2])
	fmt.Printf("Slice to last element: %v\n", slice[:len(slice)-1])

	fmt.Println("Changing slice - remove first element")

	slice = append(slice[:0], slice[1:]...)

	fmt.Println("Changed slice:", slice)

	fmt.Println("Changing slice - add last element")

	slice = append(slice, 6)

	fmt.Println("Changed slice:", slice)
}
