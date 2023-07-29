/*

Packages can contain multiple code files, 
and to simplify development, 
access control rules and package
prefixes do not apply when accessing features defined in the same package. 

Add a file named tax.go to the store folder.

*/


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

// The Price method can access the unexported calcTax method, but this method—and the type it
// applies to—is available for use only within the store package.

/* Output: 

Name: Kayak
Category: Watersports
Price: 348.75
product.Price() =  348.75
product.SetPrice(newPrice) =  361.25

*/





