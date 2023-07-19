package main

import "fmt"




func main() {
	
	// names := make([]string, 3)
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	
	names = append(names, "Hat", "Gloves")
	
	fmt.Println(names)



}

// Output
// [Kayak Lifejacket Paddle Hat Gloves]