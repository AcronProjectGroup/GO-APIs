package main

import (
	"composition/store"
	"fmt"
)

func main() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, b := range boats {
		// Go allows the fields of the nested type to be accessed in two ways:
		fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name)
	}
}

/* Output:

Conventional: Kayak Direct: Kayak
Conventional: Canoe Direct: Canoe
Conventional: Tender Direct: Tender

*/

