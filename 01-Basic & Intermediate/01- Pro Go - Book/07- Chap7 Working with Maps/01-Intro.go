package main

import "fmt"

func main() {

	products := make(map[string]float64, 10)

	products["Kayak"] = 279

	products["Lifejacket"] = 48.95

	fmt.Println("Map size:", len(products))

	fmt.Println("Price:", products["Kayak"])

	fmt.Println("Price:", products["Hat"])

	fmt.Printf("%v\n", products)


}


/* Output:

Map size: 2
Price: 279
Price: 0
map[Kayak:279 Lifejacket:48.95]

*/

