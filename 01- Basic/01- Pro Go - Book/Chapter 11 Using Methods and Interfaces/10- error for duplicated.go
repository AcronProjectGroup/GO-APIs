// The compiler will report an error if I try to define a method that duplicates an existing name/
// receiver combination, regardless of whether the remaining method parameters are different
package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}
type Supplier struct {
	name, city string
}

// ...other methods omitted for brevity...
func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
}

func (supplier *Supplier) printDetails(showName bool) {
	if showName {
		fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
	} else {
		fmt.Println("Supplier:", supplier.name)
	}
}

func main() {
	products := []*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}
	for _, p := range products {
		p.printDetails()
	}
	suppliers := []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails()
	}
}

/*  Output:

./10- error for duplicated.go:20:27: method Supplier.printDetails already declared at ./10- error for duplicated.go:16:27
./10- error for duplicated.go:30:3: undefined: newProduct
./10- error for duplicated.go:31:3: undefined: newProduct
./10- error for duplicated.go:32:3: undefined: newProduct
./10- error for duplicated.go:35:5: p.printDetails undefined (type *Product has no field or method printDetails)

*/