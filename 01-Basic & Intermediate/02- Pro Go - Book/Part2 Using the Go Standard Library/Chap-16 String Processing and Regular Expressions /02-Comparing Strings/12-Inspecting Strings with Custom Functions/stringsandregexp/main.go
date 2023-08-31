package main

import (
	"fmt"
	"strings"
	// "unicode"
)

func main() {
	description := "A boat for one person"

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	

	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))

	fmt.Println("IndexFunc:", strings.IndexFunc(description, 
		func(r rune) bool {
		return r == 'B' || r == 'b'
	}))

}

/* Output:



*/
