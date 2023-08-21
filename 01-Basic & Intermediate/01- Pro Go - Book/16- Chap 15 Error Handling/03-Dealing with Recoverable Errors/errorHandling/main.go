// produce a problematic response from the TotalPrice method.

package main

import "fmt"

func main() {
	// categories := []string{"Watersports", "Chess", "Soccer"}
	categories := []string { "Watersports", "Chess", "Running" }

	for _, cat := range categories {
		total := Products.TotalPrice(cat)
		fmt.Println(cat, "Total:", ToCurrency(total))
	}

}

/* Output:

Watersports Total: $328.95
Chess Total: $1291.00
Running Total: $0.00

*/

