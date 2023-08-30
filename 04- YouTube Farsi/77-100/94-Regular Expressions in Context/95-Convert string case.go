// Use the ToLower, ToUpper, Title, or ToTitle function in the strings package

// Converting String Case

/*

Function			Description
------------		------------------------------------------------
ToLower(str)		This function returns a new string containing the characters in the specified string
					mapped to lowercase.

ToUpper(str)		This function returns a new string containing the characters in the specified string
					mapped to lowercase.

Title(str)			This function converts the specific string so that the first character of each word is
					uppercase and the remaining characters are lowercase.

ToTitle(str)		This function returns a new string containing the characters in the specified string
					mapped to title case.

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for sailing"
	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.ToTitle(description))
}

