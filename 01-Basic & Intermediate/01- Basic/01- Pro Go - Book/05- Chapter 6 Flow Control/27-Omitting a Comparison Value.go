package main

import (
	"fmt"
	//"strconv"
)

func main() {
	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is >= 3 && < 7")
		default:
			fmt.Println(counter, "is >= 7")
		}
	}
}



/*

Zero value
1 is < 3
2 is < 3
3 is >= 3 && < 7
4 is >= 3 && < 7
5 is >= 3 && < 7
6 is >= 3 && < 7
7 is >= 7
8 is >= 7
9 is >= 7

*/