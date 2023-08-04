package main

import "fmt"

func CalcStoreTotal(data ProductData) {
	var storeTotal float64

	for category, group := range data {
		storeTotal += group.TotalPrice(category)
	}

	fmt.Println("Total:", ToCurrency(storeTotal))
}


func (group ProductGroup) TotalPrice(category string) (total float64) {

	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
	}
	
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return

}

/* Output:

main function started
Watersports product: Kayak
Watersports product: Lifejacket
Watersports subtotal: $328.95
Soccer product: Soccer Ball
Soccer product: Corner Flags
Soccer product: Stadium
Soccer subtotal: $79554.45
Chess product: Thinking Cap
Chess product: Unsteady Chair
Chess product: Bling-Bling King
Chess subtotal: $1291.00
Total: $81174.40
main function complete

*/