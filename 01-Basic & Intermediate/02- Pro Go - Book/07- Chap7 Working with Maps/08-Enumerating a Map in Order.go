package main

import (
	"fmt"
	"sort"
)

func main() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	

	// Get Lenght of Map
	keys := make([]string, 0, len(products))
	
	// Import values in index of "keys" slice
	for key, forNothing := range products {
		if forNothing == 0 {
			forNothing += 0
		} 
		keys = append(keys, key)
	}
	
	// Sort it "keys" slice
	sort.Strings(keys)
	
	// Print it "keys" slice
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", products[key])
	}
}



/* Output:

Key: Hat Value: 0
Key: Kayak Value: 279
Key: Lifejacket Value: 48.95

*/

