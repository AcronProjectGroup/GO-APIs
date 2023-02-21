// Source : https://linuxhint.com/golang-map-of-maps/

package main

import (
	"fmt"
	"reflect"
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

	// Diagnosis Type = Output:   map
	fmt.Println(reflect.ValueOf(nested).Kind())
}


