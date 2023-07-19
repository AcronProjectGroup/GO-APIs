package main

import "fmt"

func main() {
	
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	
	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Kayak"])
	fmt.Println("Price:", products["Hat"])

	fmt.Println("\n", products)
}

/* Output: 

Map size: 2
Price: 279
Price: 0

 map[Kayak:279 Lifejacket:48.95]

 */

 

