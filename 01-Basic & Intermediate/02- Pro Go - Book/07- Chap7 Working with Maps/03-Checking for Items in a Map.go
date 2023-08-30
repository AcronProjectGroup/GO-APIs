package main

import "fmt"

func main() {
	
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	
	fmt.Println("Hat:", products["Hat"])

}

/* Output: 

The problem with this code is that products["Hat"] returns zero, 
but it isn’t known whether this
is because zero is the stored value or because there is no value associated with the key Hat. 
To solve this problem, maps produce two values when reading a value

*/