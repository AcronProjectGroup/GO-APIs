/*

Verb 	Description
-----	-------------------
%s 		Prints the value as plain string
%q 		Prints the value as a double-quoted string
%8s 	Prints the value as plain string (width 8, right justified)
%-8s 	Prints the value as plain string (width 8, left justified)
%x 		Prints the value as hex dump of byte values
% x 	Prints the value as hex dump with spaces

*/

package main

import (
	"fmt"
)

func main() {
	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)
}



/*

Result:
Hello
"Hello"
   Hello
Hello
48656c6c6f
48 65 6c 6c 6f

*/