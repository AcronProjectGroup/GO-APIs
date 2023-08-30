// Care must be taken when copying structs to consider the effect on pointer fields

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


func main() {
	acme := &Supplier{"Acme Co", "New York"}
	p1 := newProduct("Kayak", "Watersports", 275, acme)

	p2 := *p1

	p1.name = "Original Kayak"
	p1.Supplier.name = "BoatCo"
	
	
	/*
	
	Output:
		fmt.Println(*p1.Supplier)
		fmt.Println(*p2.Supplier)
		{BoatCo New York}
		{BoatCo New York}
	*/

	for _, p := range []Product{*p1, p2} {
		fmt.Println(
			p,
			"Name:", p.name,"	", 
			"Supplier-name:", p.Supplier.name,"	", 
			"Suplier-sity:",p.Supplier.city,"	",
		)
	}

}



