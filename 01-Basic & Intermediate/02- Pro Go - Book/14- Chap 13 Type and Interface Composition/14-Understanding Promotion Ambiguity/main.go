// A related issue arises when two embedded fields use the same field or method names
// An Ambiguous Method

package main

import (
	"composition/store"
	"fmt"
)

func main() {
	kayak := store.NewProduct("Kayak", "Watersports", 279)
	type OfferBundle struct {
		*store.SpecialDeal
		*store.Product
	}
	
	bundle := OfferBundle{
		store.NewSpecialDeal("Weekend Special", kayak, 50),
		store.NewProduct("Lifrejacket", "Watersports", 48.95),
	}
	
	
	fmt.Println("Price:", bundle.Price(0))
	// Output runing:
		// .\main.go:22:33: ambiguous selector bundle.Price
	

}
