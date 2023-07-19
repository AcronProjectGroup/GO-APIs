package main

import (
	"fmt"
	//"strconv"
)

func main() {
	
	var price = "€48.95"
	
	for index, char := range price {
		fmt.Println((rune(char)))
		fmt.Println(index, char, string(char))
	}

}

/*

Strings.go 
0 8364 €
3 52 4
4 56 8
5 46 .
6 57 9
7 53 5


*/