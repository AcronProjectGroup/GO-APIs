package main

import "fmt"

func main() {
	
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	
	replacementProducts := []string{"Canoe", "Boots"}
	
	copy(products[0:1], replacementProducts)
	
	fmt.Println("products:", products)

}

/* Output:
products: [Canoe Lifejacket Paddle Hat]
*/