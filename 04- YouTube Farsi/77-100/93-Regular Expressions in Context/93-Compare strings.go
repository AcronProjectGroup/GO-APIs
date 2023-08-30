// Use the Contains, EqualFold, or Has* function in the strings package

// Comparing Strings

// The strings package provides comparison functions

/*

Function				Description
-----------				-------------------------------------------------------------------------
Contains(s, substr)		This function returns true if the string s contains substr and false if it does not.

ContainsAny(s, substr)	This function returns true if the string s contains any of the characters
						contained in the string substr.

						ContainsRune(s, rune)	This function returns true if the string s contains a specific rune.

EqualFold(s1, s2)		This function performs a case-insensitive comparison and returns true of
						strings s1 and s2 are the same.
HasPrefix(s, prefix)	This function returns true if the string s begins with the string prefix.

HasSuffix(s, suffix)	This function returns true if the string ends with the string suffix.


*/

package main
import (
	"fmt"
	"strings"
)
func main() {
	product := "Kayak"
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
}