package main

import "fmt"

func main() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	
	for key, value := range products {
		fmt.Printf("Key: %s	Value: %.f\n", key, value)
	}

}



/* Output:

Key: Kayak      Value: 279
Key: Lifejacket Value: 49
Key: Hat        Value: 0


*/