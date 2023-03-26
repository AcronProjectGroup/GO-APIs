package main

import "fmt"

func main() {
	var num int = 42
	var p *int = &num
	fmt.Println(*p) // Output: 42

	var userInput string = "kjsdhg1324"
	var myPointer = &userInput
	fmt.Println(*myPointer)
}
