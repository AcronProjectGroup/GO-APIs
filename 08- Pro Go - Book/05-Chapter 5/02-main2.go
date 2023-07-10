package main

import "fmt"

func main() {
	price, tax := 275.00, 27.40
	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax
	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)
}



