package main

import "fmt"

func main() {
	// first := 100
	
	var second *int
	fmt.Println(second)


	var three int // string ""
	fmt.Println(*&three)
	fmt.Println(three)

	// second = &first
	// fmt.Println(second)
}

/*
Output

<nil>
0xc000018098

*/
