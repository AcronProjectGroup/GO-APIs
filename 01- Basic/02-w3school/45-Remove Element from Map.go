// Syntax
// delete(map_name, key)

package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)

	delete(a, "year")

	fmt.Println(a)
}

// Result:
// map[brand:Ford model:Mustang year:1964]
// map[brand:Ford model:Mustang]
