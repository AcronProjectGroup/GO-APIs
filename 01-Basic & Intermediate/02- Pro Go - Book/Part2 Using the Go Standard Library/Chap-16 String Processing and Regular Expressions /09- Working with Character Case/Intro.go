/*
Function        Description
------------    ----------------------------------------------------------------------------
IsLower(rune)   This function returns true if the specified rune is lowercase.
ToLower(rune)   This function returns the lowercase rune associated with the specified rune.
IsUpper(rune)   This function returns true if the specified rune is uppercase.
ToUpper(rune)   This function returns the upper rune associated with the specified rune.
IsTitle(rune)   This function returns true if the specified rune is title case.
ToTitle(rune)   This function returns the title case rune associated with the specified rune.
*/

package main

import (
	"fmt"
	//"strings"
	"unicode"
)

func main() {
	
	product := "Kayak"
	
	for _, char := range product {
		fmt.Println(char,"	",string(char), "Upper case:", unicode.IsUpper(char))
	}

}

