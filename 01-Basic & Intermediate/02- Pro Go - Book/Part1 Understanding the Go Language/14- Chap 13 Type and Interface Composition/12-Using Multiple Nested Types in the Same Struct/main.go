package main

import (
	"composition/store"
	"fmt"
)

func main() {
	
	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true,"Dora", "Charlie"),
	}
	
	for _, r := range rentals {
		fmt.Println(
			"Rental Boat:", r.Name, 
			"Rental Price:", r.Price(0.2),
			"Captain:", r.Captain)
	}

}

/* Output: 

Rental Boat: Rubber Ring Rental Price: 12 Captain: N/A
Rental Boat: Yacht Rental Price: 60000 Captain: Bob
Rental Boat: Super Yacht Rental Price: 120000 Captain: Dora

*/
