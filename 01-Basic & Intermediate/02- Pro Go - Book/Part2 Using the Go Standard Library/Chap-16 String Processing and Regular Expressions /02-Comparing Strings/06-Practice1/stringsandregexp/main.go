package main

import (
	"fmt"
	"strings"
)
// strings.Title is deprecated: The rule Title uses for word boundaries does not handle Unicode punctuation properly. Use golang.org/x/text/cases instead.deprecated(default)


func main() {
	description := "A boat for sailing"
	
	fmt.Println("Original:", description)
	fmt.Println("ToUpper:", strings.Title(description))
	fmt.Println("ToLower:", strings.ToLower(description))
	fmt.Println("ToTitle:", strings.ToTitle(description))
	fmt.Println("ToUpper:", strings.ToUpper(description))

}

/*
Output:

Original: A boat for sailing
ToLower: a boat for sailing
ToTitle: A BOAT FOR SAILING
ToUpper: A BOAT FOR SAILING


*/