// Dependencies on custom packages are declared using the import statement

// package main

// import (
// 	"fmt"
// 	"UsingCustomPackage05/packages/store"
// )

// func main() {
// 	product := store.Product{
// 		Name:     "Kayak",
// 		Category: "Watersports",
// 	}
// 	fmt.Println("Name:", product.Name)
// 	fmt.Println("Category:", product.Category)
// }

package main

import (
	"fmt"
	currencyFmt "UsingCustomPackage05/packages/fmt"

	"UsingCustomPackage05/packages/store"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))
}
