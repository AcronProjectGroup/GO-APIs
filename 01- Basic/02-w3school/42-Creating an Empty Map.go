// Syntax
// var a map[KeyType]ValueType

// Note: The make()function is the right way to create an empty map.
// If you make an empty map in a different way and write to it,
// it will causes a runtime panic.

package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)
}

// Result:
// false
// true

// Allowed Key Types

// The map key can be of any data type for which the equality operator (==) is defined. These include:

//     Booleans
//     Numbers
//     Strings
//     Arrays
//     Pointers
//     Structs
//     Interfaces (as long as the dynamic type supports equality)

// Invalid key types are:

//     Slices
//     Maps
//     Functions

// These types are invalid because the equality operator (==) is not defined for them.
