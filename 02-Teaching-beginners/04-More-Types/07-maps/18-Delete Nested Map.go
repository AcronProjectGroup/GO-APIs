// Source : https://linuxhint.com/golang-map-of-maps/

package main

import (
	"fmt"
	// "reflect"
)


func main() {
    nested := map[int]map[string]string{
        1: {
            "a": "Apple",
            "b": "Banana",
            "c": "Coconut",
        },
        2: {
            "a": "Tea",
            "b": "Coffee",
            "c": "Milk",
        },
        3: {
            "a": "Italian Food",
            "b": "Indian Food",
            "c": "Chinese Food",
        },
    }

	// for k, v := range nested {
	// 	fmt.Println("Key:", k, "Value:", v)
		// fmt.Println(reflect.ValueOf(k).Kind())
		// fmt.Println(reflect.ValueOf(v).Kind())
	// }

	// To iterate over an individual map
	// for k, v := range nested[2] {
	// 	fmt.Println(k, v)
	// }


	// Diagnosis Type = Output:   map
	// fmt.Println(reflect.ValueOf(nested).Kind())

	delete(nested[1], "a")
	fmt.Println(nested)

	for i := 1; i < 4; i++ {
		delete(nested[i], "a")
		delete(nested[i], "a")
		delete(nested[i], "a")
		fmt.Println(nested)
	}



}


