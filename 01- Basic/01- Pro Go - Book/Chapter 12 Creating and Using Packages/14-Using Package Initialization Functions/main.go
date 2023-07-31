package main

import (
	. "UsingCustomPackage05/packages/fmt"
	"UsingCustomPackage05/packages/store"
	"UsingCustomPackage05/packages/store/cart"
	"fmt"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}


	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", ToCurrency(cart.GetTotal()))


}

/* Output: 

Name: Alice
Total: $348.75

*/
