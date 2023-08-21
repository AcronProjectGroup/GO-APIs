package main

import ( //init GetData
	"fmt"
	_ "UsingCustomPackage05/packages/data"
	. "UsingCustomPackage05/packages/fmt"
	"UsingCustomPackage05/packages/store"
	"UsingCustomPackage05/packages/store/cart"
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

data.go init function invoked
data.go init function invoked
data.go init function invoked
data.go init function invoked
Name: Alice
Total: $300.00

*/