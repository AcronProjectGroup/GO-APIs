package main

import "fmt"



func printPrice(string, float64, float64) {
	// taxAmount := price * 0.25
	fmt.Println("No parameters")
}


func main() {
	printPrice("Kayak", 275, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice("Soccer Ball", 19.50, 0.15)
}


/*

Parameters without names cannot be accessed within the function, 
and this feature is mainly used in conjunction with interfaces


*/