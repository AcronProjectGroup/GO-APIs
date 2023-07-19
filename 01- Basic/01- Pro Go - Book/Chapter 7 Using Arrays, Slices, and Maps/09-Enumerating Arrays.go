package main

import "fmt"

func main() {
	
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	
	for index, value := range names {
	
		fmt.Println("Index:", index, "Value:", value)
	}
	
}
