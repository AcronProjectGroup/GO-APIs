package main

import "fmt"

func main() {
	categories := []string{"Watersports", "Chess", "Soccer"}

	for _, cat := range categories {
		total := Products.TotalPrice(cat)
		fmt.Println(cat, "Total:", ToCurrency(total))
	}

}


