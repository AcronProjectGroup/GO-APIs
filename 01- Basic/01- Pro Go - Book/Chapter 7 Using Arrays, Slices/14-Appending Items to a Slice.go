package main

import "fmt"

func main() {
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	
	appendedNames := append(names, "Hat", "Gloves")
	
	names[0] = "Canoe"
	
	
	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)

}


/*

Output:
names: [Canoe Lifejacket Paddle]
appendedNames: [Kayak Lifejacket Paddle Hat Gloves]

*/