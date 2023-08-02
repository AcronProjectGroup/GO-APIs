// Go uses interfaces to describe methods that can be implemented by
// multiple types.





// Using an Interface
package main

import (
	"composition/store"
	"fmt"
)

func main() {
	
	
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}




	for key, p := range products {
		fmt.Println("Key:", key, "Price:", p.Price(0.2))
	}


}
