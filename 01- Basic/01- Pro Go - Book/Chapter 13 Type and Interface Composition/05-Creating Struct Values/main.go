package main

import (
	"composition/store"
	"fmt"
)



func main() {
	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}
	

	for _, p := range []*store.Product{kayak, lifejacket} {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}

}

/* Output: 

Name: Kayak Category: Watersports Price: 330
Name: Lifejacket Category: Watersports Price: 0

*/
