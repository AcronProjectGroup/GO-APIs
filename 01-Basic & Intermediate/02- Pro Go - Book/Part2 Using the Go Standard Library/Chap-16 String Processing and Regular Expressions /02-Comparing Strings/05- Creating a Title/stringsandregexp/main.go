package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for sailing"
	
	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.Title(description))

}

/*
Original: A boat for sailing
Title: A Boat For Sailing
*/