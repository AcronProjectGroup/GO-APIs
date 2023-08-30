/*

External packages are imported and used like custom packages. 

The import statement specifies the module path, 
and the last part of that path is used to access 
the features exported by the package. 

In this case, the package is named color, 
and this is the prefix used to access the package features.
The Green and Cyan functions used.


*/


package main

import (
	//"fmt"
	_ "UsingCustomPackage05/packages/data"
	. "UsingCustomPackage05/packages/fmt"
	"UsingCustomPackage05/packages/store"
	"UsingCustomPackage05/packages/store/cart"
	"github.com/fatih/color" 
)


func main() {

	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}

	color.Green("\nName: " + cart.CustomerName)
	color.Cyan("\nTotal: " + ToCurrency(cart.GetTotal()))
	color.Blue("\nSINA LALEHBAKHSH")
	color.BlueString("\nThis is Blue String")
	println()


}
