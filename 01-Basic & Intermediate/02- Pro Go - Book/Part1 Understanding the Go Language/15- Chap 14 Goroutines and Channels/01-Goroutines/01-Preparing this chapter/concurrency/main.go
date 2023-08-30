package main

import "fmt"

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	fmt.Println("main function complete")
}




/*

go run main.go 
Ouput:
	# command-line-arguments
	./main.go:7:2: undefined: CalcStoreTotal
	./main.go:7:17: undefined: Products

go run .
Output:
	main function started
	Watersports subtotal: $328.95
	Soccer subtotal: $79554.45
	Chess subtotal: $1291.00
	Total: $81174.40
	main function complete

*/