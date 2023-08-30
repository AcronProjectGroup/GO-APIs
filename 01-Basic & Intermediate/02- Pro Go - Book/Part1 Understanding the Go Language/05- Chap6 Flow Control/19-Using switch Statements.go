package main

import (
	"fmt"
	//"strconv"
)

func main() {
	product := "Kayak"
	for index, character := range product {
		switch character {
		case 'K':
			fmt.Println("K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}
