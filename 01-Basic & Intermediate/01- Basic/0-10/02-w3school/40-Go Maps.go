// Go Maps

// Maps are used to store data values in key:value pairs.

// Each element in a map is a key:value pair.

// A map is an unordered and changeable collection that does not allow duplicates.

// The length of a map is the number of its elements. You can find it using the len() function.

// The default value of a map is nil.

// Maps hold references to an underlying hash table.

// Go has multiple ways for creating maps.

// Creating Maps Using var and :=
// Syntax
// var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}

package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}

// Result:
// a   map[brand:Ford model:Mustang year:1964]
// b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]
