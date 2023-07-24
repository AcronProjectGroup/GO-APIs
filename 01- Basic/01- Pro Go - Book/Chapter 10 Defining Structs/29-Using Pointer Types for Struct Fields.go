// Pointers can also be used for struct fields, including pointers to other struct types

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

	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275, acme),
		newProduct("Hat", "Skiing", 42.50, acme),
	}

	for _, p := range products {
		fmt.Println(
			"Name:", p.name, "	",
			"Supplier-name:", p.Supplier.name, "	",
			"Suplier-city:", p.Supplier.city, "	",
		)
	}

}
/* Output:

Name: Kayak      Supplier-name: Acme Co          Suplier-city: New York 
Name: Hat        Supplier-name: Acme Co          Suplier-city: New York

*/