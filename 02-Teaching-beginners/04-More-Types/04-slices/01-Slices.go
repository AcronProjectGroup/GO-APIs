package main

import "fmt"

func main() {

	primes := []int{}
	primes = append(primes, 200, 300, 400)
	fmt.Println(primes)
	// --------------------------------------
	primes = []int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
	// --------------------------------------
	var list []int 
	list = append(list, 1, 2, 3 )
	fmt.Println(list)

}
