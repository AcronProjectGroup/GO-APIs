package main

import (
	"fmt"
	"strings"
)

func main() {
	product := "AcronProject"

	fmt.Println("Contains:", strings.Contains(product, "ron"))
	
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'P'))
	
	fmt.Println()
	fmt.Println("EqualFold:", strings.EqualFold(product, "ACRONPROJECT"))
	fmt.Println("EqualFold:", strings.EqualFold(product, "acronproject"))
	fmt.Println()
	
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ac"))
	
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "ct"))
}
