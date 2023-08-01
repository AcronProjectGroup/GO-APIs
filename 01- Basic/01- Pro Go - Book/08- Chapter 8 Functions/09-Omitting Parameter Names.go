package main

import "fmt"

func printPrice(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}



func main() {
	printPrice("Kayak", 275, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice("Soccer Ball", 19.50, 0.15)
}




/*

The underscore is known as the blank identifier, 
and the result is a parameter for which a value must be
provided when the function is called but whose value 
cannot be accessed inside the function’s code block.
This may seem like an odd feature, 
but it can be a useful way to indicate that a parameter is not used within a
function, which can arise when implementing the methods required by an interface.


*/