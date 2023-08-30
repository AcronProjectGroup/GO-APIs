// The defer keyword is used to schedule a function call that will be performed immediately before the current function returns

package main

import "fmt"

func calcTotalPrice(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	
	fmt.Println("Function about to return")

	return

}



func main() {

	defer fmt.Println("Sina")
	defer fmt.Println("3333333")

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	_, total := calcTotalPrice(products)
	fmt.Println("Total:", total)

}


/* Output:

Keyword.go 
Function started
Function about to return
Second defer call
First defer call
Total: 323.95
3333333
Sina

*/


