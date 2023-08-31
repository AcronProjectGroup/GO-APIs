package main

import (
	"fmt"
	"strings"
)

// strings.Title is deprecated: The rule Title uses for word boundaries does not handle Unicode punctuation properly. Use golang.org/x/text/cases instead.deprecated(default)

func main() {

	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))
	

}

/*
Output:



*/
