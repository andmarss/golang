package main

import "fmt"

var car = "mercedes"

func main(){
	checkSpareWheel()
}

func checkSpareWheel(){
	if car == "ford" {
		fmt.Println("spare wheel exist")
	} else {
		fmt.Println("spare wheel not exist")
	}
}