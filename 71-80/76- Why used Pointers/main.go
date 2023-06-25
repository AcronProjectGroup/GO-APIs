package main

import (
	"fmt"
)


func zero(x int) {
	x = 0
}


func zero123(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5

	// -----------------
	xCanBeChanges := 5
	zero123(&xCanBeChanges)
	fmt.Println(xCanBeChanges) // x is 0

}



