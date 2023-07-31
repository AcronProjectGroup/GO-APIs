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
