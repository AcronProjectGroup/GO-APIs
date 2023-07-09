package main

import "fmt"

func main() {
	first := 100
	
	second := &first
	
	first++
	
	*second++
	
	var myNewPointer *int
	

	// fmt.Println(myNewPointer)
	// fmt.Println(&myNewPointer)
	// fmt.Println(*myNewPointer) // Error
	

	// myNewPointer = *&second

	
	myNewPointer = second
	
	*myNewPointer++
	
	fmt.Println("First:", first)
	
	fmt.Println("Second:", *second)
}
