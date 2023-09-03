package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for one person"
	
	trimmed := strings.Trim(description, "Asno ")
	
	fmt.Println("Trimmed:", trimmed)

}

/*

Output:

Trimmed: boat for one per


*/