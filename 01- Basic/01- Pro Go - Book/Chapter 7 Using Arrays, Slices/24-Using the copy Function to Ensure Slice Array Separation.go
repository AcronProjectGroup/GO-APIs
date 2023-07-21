package main

import "fmt"

func main() {
	
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	
	allNames := products[1:]
	
	someNames := make([]string, 2)
	
	copy(someNames, allNames)
	
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)


}


/* Output:
someNames: [Lifejacket Paddle]
allNames [Lifejacket Paddle Hat]
*/