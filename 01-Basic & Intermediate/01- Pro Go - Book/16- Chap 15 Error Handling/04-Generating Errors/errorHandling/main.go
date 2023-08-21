// The outcome from invoking the TotalPrice method is determined 
// by examining the combination of the two results.

package main

import "fmt"

func main() {
	categories := []string{"Watersports", "Chess", "Running"}

	for _, cat := range categories {
		// total := Products.TotalPrice(cat)
		// fmt.Println(cat, "Total:", ToCurrency(total))
		total, err := Products.TotalPrice(cat)
		if err == nil {
			fmt.Println(cat, "Total:", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}

}

/* Output:

Watersports Total: $328.95
Chess Total: $1291.00
Running (no such category)

*/
