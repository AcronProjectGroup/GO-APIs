package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price - 10, supplier}
}

/*
To ensure the Supplier is duplicated, 
the copyProduct function assigns it to a separate variable and
then creates a pointer to that variable. 
This is awkward, but the effect is to force a copy of the struct, albeit
this is a technique that is specific to a single struct type 
and must be repeated for each nested struct field.
*/
func copyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}


func main() {
	acme := &Supplier{"Acme Co", "New York"}
	p1 := newProduct("Kayak", "Watersports", 275, acme)
	p2 := copyProduct(p1)
	
	p1.name = "Original Kayak"
	p1.Supplier.name = "BoatCo"

	fmt.Println(*p1)
	fmt.Println(p2)

	for _, p := range []Product{*p1, p2} {
		fmt.Println("Name:", p.name, "Supplier:",
			p.Supplier.name, p.Supplier.city)
	}
}
