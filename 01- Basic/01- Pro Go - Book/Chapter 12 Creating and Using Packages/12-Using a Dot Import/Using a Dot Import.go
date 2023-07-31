/*

There is a special alias, known as the dot import,
that allows a package’s features to be used without using a prefix.

When using a dot import, you must ensure that 
the names of the features imported from the package
are not defined in the importing package. 

For the example, this means I must ensure that the name
ToCurrency is not used by any feature defined in the main package. 
For this reason, dot imports should be used with caution.


*/

package main

import (
	"fmt"

	// A dot import uses a period as the package alias
	// The dot import allows me to access the ToCurrency function without using a prefix
	. "UsingCustomPackage05/packages/fmt"
	
	
	"UsingCustomPackage05/packages/store"
)



func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", ToCurrency(product.Price()))
}
