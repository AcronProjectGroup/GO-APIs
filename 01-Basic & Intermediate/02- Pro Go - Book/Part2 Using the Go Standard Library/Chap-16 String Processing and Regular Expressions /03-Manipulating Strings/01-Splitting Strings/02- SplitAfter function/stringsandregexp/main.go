/*
Function				Description
---------				----------------------------------------------------------------------
Fields(s)				This function splits a string on whitespace characters and returns a slice
						containing the nonwhitespace sections of the string s.

FieldsFunc(s, func)		This function splits the string s on the characters for which a custom function
						returns true and returns a slice containing the remaining sections of the string.

Split(s, sub)			This function splits the string s on every occurrence of the specified substring,
						returning a string slice. If the separator is the empty string, then the slice will
						contain strings for each character.

SplitN(s, sub, max)		This function is similar to Split, but accepts an additional int argument that
						specifies the maximum number of substrings to return. The last substring in the
						result slice will contain the unsplit portion of the source string.

SplitAfter(s, sub)		This function is similar to Split but includes the substring used in the results.
						See the text after the table for a demonstration.

SplitAfterN(s, sub,max)	This function is similar to SplitAfter, but accepts an additional int argument
						that specifies the maximum number of substrings to return.

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for one person"
	
	splits := strings.Split(description, " ")
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	
	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}
}

/* Output:

Split >>A<<
Split >>boat<<
Split >>for<<
Split >>one<<
Split >>person<<
SplitAfter >>A <<
SplitAfter >>boat <<
SplitAfter >>for <<
SplitAfter >>one <<
SplitAfter >>person<<


*/
