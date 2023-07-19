package main

import "fmt"

func main() {
	
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	
	replacementProducts := []string{"1", "222"}
	
	copy(products, replacementProducts)
	
	fmt.Println("products:", products)

}


/* Output:
products: [Canoe Boots Paddle Hat]

*/
