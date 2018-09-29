package main

import "fmt"

/*
1. Declaring an array
var arrayName [length]type

2. Setting a value in the array
arrayName[index] = value

3. Getting a value in the array
arrayName[index]
*/

func main() {
	var arr [5]int
	fmt.Printf("Contents of array: %v\n", arr)

	populateIntegerArray(arr)
	fmt.Printf("Contents of array: %v\n", arr)

	arr = fillIntegerArray(arr)
	fmt.Printf("Contents of array: %v\n", arr)

	fmt.Println("Length of the array:", len(arr))

	matrix := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	fmt.Println(matrix)

	printThreeByFourMatrtix(matrix)
}

func populateIntegerArray(arr [5]int) {
	arr[0] = 3
	arr[1] = 6
	arr[2] = 9
	arr[3] = 12
	arr[4] = 15
}

func fillIntegerArray(arr [5]int) [5]int {
	arr[0] = 3
	arr[1] = 6
	arr[2] = 9
	arr[3] = 12
	arr[4] = 15

	return arr
}

func printThreeByFourMatrtix(arr [3][4]int) {
	rowLength := len(arr)
	columnLength := len(arr[0])

	for i := 0; i < rowLength; i++ {
		for j := 0; j < columnLength; j++ {
			fmt.Printf("%5d", arr[i][j])
		}
		fmt.Println()
	}
}
