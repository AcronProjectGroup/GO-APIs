package main

import (
	"fmt"
)

func main() {
	var language string = "golang"
	var pointer *string = &language

	fmt.Println("language =", language)    // output->      language = golang
	fmt.Println("pointer =", pointer)      // output->      pointer = 0xc000010250
	fmt.Println("*pointer =", *pointer)    // output->      *pointer = golang
}