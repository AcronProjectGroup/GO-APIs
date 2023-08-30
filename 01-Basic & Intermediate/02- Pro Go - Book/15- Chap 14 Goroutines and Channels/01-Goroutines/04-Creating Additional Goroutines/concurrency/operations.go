package main

import "fmt"

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	for category, group := range data {
		// Go routines make it easy to invoke functions and methods, but the change
		// has introduced a common problem
		go group.TotalPrice(category)
		// storeTotal += group.TotalPrice(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}


func (group ProductGroup) TotalPrice(category string) (total float64) {
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return
}

/* 


ubuntu@ubuntu:~/Desktop/Name/Name/concurrency$ go run .
	main function started
	Total: $0.00
	main function complete













// ubuntu@ubuntu:~/Desktop/Name/Name/concurrency$ go run main.go 
// 	# command-line-arguments
// 	./main.go:7:2: undefined: CalcStoreTotal
// 	./main.go:7:17: undefined: Products


*/