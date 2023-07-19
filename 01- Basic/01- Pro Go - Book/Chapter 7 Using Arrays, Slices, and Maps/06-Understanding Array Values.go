package main

import "fmt"

func main() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	

	// This is an image of "names" from that time  
	otherArray := names
	
	// But now when "names" changes, dont effect on "otherArray"
	names[0] = "Canoe"
	
	fmt.Println("names:", names)

	fmt.Println("otherArray:", otherArray)

}

// Output:
// names: [Canoe Lifejacket Paddle]
// otherArray: [Kayak Lifejacket Paddle]


