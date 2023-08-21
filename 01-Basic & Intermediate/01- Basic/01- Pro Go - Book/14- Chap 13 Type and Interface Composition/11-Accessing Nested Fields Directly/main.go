package main

import (
	"composition/store"
	"fmt"
)

func main() {
	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false),
		store.NewRentalBoat("Yacht", 50000, 5, true, true),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true),
	}


	// Go promotes fields from the nested Boat and Product types 
	// so they can be accessed through the top-
	// level RentalBoat type, which allows the Name field to be read
	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2))
	}
	
}

/* Output:

Rental Boat: Rubber Ring Rental Price: 12
Rental Boat: Yacht Rental Price: 60000
Rental Boat: Super Yacht Rental Price: 120000

*/