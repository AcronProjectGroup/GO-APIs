package main

import "fmt"

func main() {
	first := 100
	
	var second *int
	fmt.Println(second)

	second = &first
	fmt.Println(second)
}

/*
Output

<nil>
0xc000018098

*/
