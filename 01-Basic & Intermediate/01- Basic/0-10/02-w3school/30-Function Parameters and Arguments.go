// Syntax
// func FunctionName(param1 type, param2 type, param3 type) {
//   // code to be executed
// }

// Note: When a parameter is passed to the function,
// it is called an argument.
// So, from the example above:
// fname is a parameter, while Liam, Jenny and Anja are arguments.

package main

import (
	"fmt"
)

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func main() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
}

// Result:
// Hello Liam Refsnes
// Hello Jenny Refsnes
// Hello Anja Refsnes

