package main

import (
	"fmt"
	//"strconv"
)

func main() {
	product := "Kayak"
	for index := range product {
		fmt.Println("Index:", index)
	}
}
