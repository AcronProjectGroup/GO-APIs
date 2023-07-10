package main

import "fmt"

func main() {
	first := 100   // Location 1 in RAM
	
	
	second := first // Location 2 in RAM
	
	var third int = first
	
	first++ // 101 in Location 1 value is changes
	
	fmt.Println("first:", first)
	
	fmt.Println("second:", second)

	fmt.Println("third:",third)
}



// Output:
	// First: 101
	// Second: 100



