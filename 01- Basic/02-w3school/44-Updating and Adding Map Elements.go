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

	a["year"] = "1970" // Updating an element
	a["color"] = "red" // Adding an element

	fmt.Println(a)
}

// Result:
// map[brand:Ford model:Mustang year:1964]
// map[brand:Ford color:red model:Mustang year:1970]
