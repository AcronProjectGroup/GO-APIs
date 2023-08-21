// Omitting the Parameter Data Type in the main.go File in the functions Folder

package main

import "fmt"

func printPrice(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}



func main() {
	printPrice("Kayak", 275, 0.2)
	printPrice("Lifejacket", 48.95, 0.2)
	printPrice("Soccer Ball", 19.50, 0.15)
}









/*

The price and taxRate parameters are both float64,
and since they are adjacent, the data type is
applied only to the final parameter of that type.

*/
