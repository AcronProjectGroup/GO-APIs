package main

import (
	"UsingCustomPackage05/packages/store"
	"fmt"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())


	fmt.Println("product.Price() = ", product.Price())


	var newPrice float64 = 10
	product.SetPrice(newPrice)	
	fmt.Println("product.SetPrice(newPrice) = ",product.Price())

}

/* Output:

Name: Kayak
Category: Watersports
Price: 279
product.Price() =  279
product.SetPrice(newPrice) =  289


*/