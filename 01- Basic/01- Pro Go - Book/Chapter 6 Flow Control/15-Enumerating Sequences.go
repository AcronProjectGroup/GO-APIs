package main

import (
	"fmt"
	//"strconv"
)

func main() {
	product := "Kayak"
	for index, character := range product {
		fmt.Println("Index:", index, "Character:", string(character))
	}
}


