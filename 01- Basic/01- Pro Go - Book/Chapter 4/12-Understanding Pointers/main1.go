package main

import "fmt"

func main() {
	first := 100   // Location 1 in RAM
	
	
	second := first // Location 2 in RAM
	

	first++ // in Location 1 value is changes
	
	fmt.Println("First:", first)
	
	fmt.Println("Second:", second)
}



// Output:
	// First: 101
	// Second: 100



