// Note Go does not support optional parameters or default values for parameters.


package main

import "fmt"



func printPrice(product string, price float64, taxRate float64) {

	taxAmount := price * taxRate

	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}




func main() {
	printPrice("Kayak", 275, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice("Soccer Ball", 19.50, 0.15)
}


/* Output:

Kayak price: 275 Tax: 55
Lifejacket price: 48.95 Tax: 9.790000000000001
Soccer Ball price: 19.5 Tax: 2.925


*/




/*

Values for parameters are supplied as arguments when invoking the function, 
meaning that different
values can be provided each time the function is called.


*/