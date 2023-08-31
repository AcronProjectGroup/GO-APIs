/*

Function				Description
---------				-----------------------------------------------------
Count(s, sub)			This function returns an int that reports how many times the specified
						substring is found in the string s.

Index(s, sub)			These functions return the index of the first or last occurrence of a specified
LastIndex(s, sub)		substring string within the string s, or -1 if there is no occurrence.
						
IndexAny(s, chars)		These functions return the first or last occurrence of any character in the
LastIndexAny(s, chars)	specified string within the string s, or -1 if there is no occurrence.
						
IndexByte(s, b)			These functions return the index of the first or last occurrence of a specified
LastIndexByte(s, b)		byte within the string s, or -1 if there is no occurrence.


IndexFunc(s, func)		These functions return the index of the first or last occurrence of the
LastIndexFunc(s, func)	character in the string s for which the specified function returns true, as
						described in the “Inspecting Strings with Custom Functions” section.

*/
package main

import (
	"fmt"
	"strings"
	// "unicode"
)

func main() {
	description := "A boat for one person"
	
	fmt.Println("Count:", strings.Count(description, "o"))
	fmt.Println("Index:", strings.Index(description, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))
}

/* Output:
Count: 4
Index: 3
LastIndex: 19
IndexAny: 2
LastIndex: 19
LastIndexAny: 4
*/

