package main

import "fmt"

func main() {
	
	names := make([]string, 3, 6)
	
	names[0] = "Kayak"
	
	names[1] = "Lifejacket"
	
	names[2] = "Paddle"
	
	moreNames := []string{"Hat Gloves"}
	
	appendedNames := append(names, moreNames...)
	
	fmt.Println("appendedNames:", appendedNames)
}

// appendedNames: [Kayak Lifejacket Paddle Hat Gloves]