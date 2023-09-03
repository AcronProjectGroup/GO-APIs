package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat AA for nn one person"

	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	trimmed := strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed:", trimmed)
}

/*

Output:

Trimmed:  boat AA for nn one perso


*/
