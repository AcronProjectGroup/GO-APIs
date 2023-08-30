package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price          float64
	}
	
	
	kayak := Product{
		name:     "Kayak",
		category: "Watersports",
	}
	
	fmt.Println(kayak.name, kayak.category, kayak.price)
	
	
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)
	
	
	var lifejacket Product
	fmt.Println("Name is zero value:", lifejacket.name == "")
	fmt.Println("Category is zero value:", lifejacket.category == "")
	fmt.Println("Price is zero value:", lifejacket.price == 0)
}

/*

The type of the lifejacket variable is Product, but no values are assigned to its fields.
The value of all the lifejacket fields is the zero value for their type.

*/




/* Output:

Kayak Watersports 0
Changed price: 300
Name is zero value: true
Category is zero value: true
Price is zero value: true

*/