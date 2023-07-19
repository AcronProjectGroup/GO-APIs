package main

import (
	"fmt"
	//"sort"
)

func main() {
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1


	fmt.Println("Equal:", p1 == nil) 
// if :
	// fmt.Println("Equal:", p1 == p2) 
	// Error arise:
	// .\main.go:13:30: invalid operation: p1 == p2 (slice can only be compared to nil)



	print(p2)

	
}
