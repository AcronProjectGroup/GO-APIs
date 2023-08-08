// Go can perform promotion only if there is no field
// or method defined with the same name on the enclosing
// type, which can lead to unexpected results.

package main

import (
	"composition/store"
	"fmt"
)

func main() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	
	Name, price, Price := deal.GetDetails()
	
	
	
	// return deal.Name, deal.price, deal.Product.Price(0)
	fmt.Println("Price method:", Price)
	
	
	fmt.Println("Name:", Name)
	fmt.Println("Price field:", price)
}
/* Output:
Price method: 279
Name: Weekend Special
Price field: 229
*/



/* In file specialdeal.go: 


// Defining a Method
func (deal *SpecialDeal) Price(taxRate float64) float64 {
	return deal.price
}

Then Output:
	Price method: 229
	Name: Weekend Special
	Price field: 229


*/