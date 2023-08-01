package main

import "fmt"

func main() {

	names := [3]string{"Kayak", "Lifejacket", "Paddle"}

	for _, value := range names {
		fmt.Println("Value:", value)
	}

}

/*

Output:
Value: Kayak
Value: Lifejacket
Value: Paddle

*/


