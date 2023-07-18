package main

import (
	"fmt"
	//"strconv"
)

func main() {
	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element:", element)
	}
}
