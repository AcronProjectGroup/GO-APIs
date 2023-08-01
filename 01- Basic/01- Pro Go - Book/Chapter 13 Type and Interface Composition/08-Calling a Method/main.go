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


	// If the field type is a value, such as Product, 
	// then any methods defined with Product or *Product
	// receivers will be promoted. If the field type is a pointer, 
	// such as *Product, then only methods with *Product
	// receivers will be prompted.
	for _, b := range boats {
		fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
	}

}



/* Output: 

Boat: Kayak Price: 330
Boat: Canoe Price: 480
Boat: Tender Price: 780.3

*/