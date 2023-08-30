// Maps are references to hash tables.

// If two map variables refer to the same hash table,
// changing the content of one variable affect the content of the other.
package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)
}

// Result:
// map[brand:Ford model:Mustang year:1964]
// map[brand:Ford model:Mustang year:1964]
// After change to b:
// map[brand:Ford model:Mustang year:1970]
// map[brand:Ford model:Mustang year:1970]
