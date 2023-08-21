package main

import (
	"fmt"
	//"strconv"
)

func main() {
	var price = "€48.95"
	

	for index, char := range []byte(price) {
	
		fmt.Println(index, char)
	
	}
}

/* Output:

0 226
1 130
2 172
3 52
4 56
5 46
6 57
7 53

*/