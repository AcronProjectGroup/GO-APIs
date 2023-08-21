package main

import "fmt"

func main() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names)
	
	// var otherArray [5]string = names
	// .\main.go:9:9: cannot use names (type [3]string) as type [4]string in assignment
	


	
	var otherArray [3]string = names // This is true
	fmt.Println(otherArray)
}
