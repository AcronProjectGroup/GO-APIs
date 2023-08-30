package main

import (
	"fmt"
	"sort"
)

func main() {
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	sort.Strings(products)

	for index, value := range products {
		fmt.Println("Index:", index, "Value:", value)
	}

}

/* Output:

Index: 0 Value: Hat
Index: 1 Value: Kayak
Index: 2 Value: Lifejacket
Index: 3 Value: Paddle

*/

