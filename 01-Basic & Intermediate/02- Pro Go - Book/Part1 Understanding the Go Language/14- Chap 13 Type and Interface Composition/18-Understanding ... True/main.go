// Accessing the Underlying Type in the main.go
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
		switch item := p.(type) {
		case *store.Product:
			fmt.Println(
				"Name:", item.Name,
				"Category:", item.Category,
				"Price:", item.Price(0.2),
			)
		case *store.Boat:
			fmt.Println(
				"Name:", item.Name,
				"Category:", item.Category,
				"Price:", item.Price(0.2),
			)
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

}

/* Output: Passed

Name: Kayak Category: Watersports Price: 334.8
Name: Soccer Ball Category: Soccer Price: 23.4
*/
