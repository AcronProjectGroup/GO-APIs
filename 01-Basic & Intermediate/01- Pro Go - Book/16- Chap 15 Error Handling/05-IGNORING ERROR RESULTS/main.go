package main

import "fmt"

func main() {
	categories := []string{"Watersports", "Chess", "Running"}
	for _, cat := range categories {
		total, _ := Products.TotalPrice(cat)
		fmt.Println(cat, "Total:", ToCurrency(total))
	}
}
