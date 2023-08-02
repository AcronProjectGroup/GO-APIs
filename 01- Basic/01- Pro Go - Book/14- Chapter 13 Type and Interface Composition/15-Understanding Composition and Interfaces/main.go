// Mixing Types

package main

import (
	"composition/store"
	"fmt"
)

func main() {
	products := map[string]*store.Product{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	
	for _, p := range products {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}

}

/* Output: 

# command-line-arguments
./main.go:12:12: cannot use store.NewBoat("Kayak", 279, 1, false) (value of type *store.Boat) as *store.Product value in map literal

*/
