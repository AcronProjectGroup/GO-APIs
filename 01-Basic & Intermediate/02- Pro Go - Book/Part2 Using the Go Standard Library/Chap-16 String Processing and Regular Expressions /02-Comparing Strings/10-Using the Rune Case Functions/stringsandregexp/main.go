package main

import (
	"fmt"
	//"strings"
	"unicode"
)

func main() {
	
	product := "Kayak"
	
	for _, char := range product {
		fmt.Println(char,"	",string(char), "Upper case:", unicode.IsUpper(char))
	}

}
