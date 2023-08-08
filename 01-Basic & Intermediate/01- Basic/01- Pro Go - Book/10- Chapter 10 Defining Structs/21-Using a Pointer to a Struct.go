package main

import "fmt"

func main() {
	
	type Product struct {
		name, category string
		price          float64
	}
	
	p1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	
	p2 := &p1
	
	// When change anything in p1 is affect on p2 so
	p1.name = "Original Kayak"
	
	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", (*p2).name)
}

/*

I used an ampersand to create a pointer to the p1 variable and assigned the address to p2, whose type
becomes *Product, meaning a pointer to a Product value. Notice that I have to use parentheses to follow the
pointer to the struct value and then read the value of the name field

The effect is that the change made to the name field is read through both p1 and p2, producing the
following output when the code is compiled and executed:
Output:

P1: Original Kayak
P2: Original Kayak

*/