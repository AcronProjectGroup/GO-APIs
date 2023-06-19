package main

import (
	"fmt"
)

func main() {

	mySlice_in_map := make(map[string][]string)

	mySlice_in_map = map[string][]string{
		"avalivn Slice": {"S yek","S do"},
		"dovomin Slice": {"S yek","S do"},
	}

	// This is shows values:
		fmt.Println(mySlice_in_map)

	// This is just point to memory place not values:
		println(mySlice_in_map)
}