package main

import "fmt"

func main() {
	first := 100

	var second = &first


	first++

	fmt.Println(first)
	fmt.Println(*second)

	*second++



	fmt.Println(first)
	fmt.Println(*second)

	
	// fmt.Println("first:", first)

	// fmt.Println("second:", *second)


}
