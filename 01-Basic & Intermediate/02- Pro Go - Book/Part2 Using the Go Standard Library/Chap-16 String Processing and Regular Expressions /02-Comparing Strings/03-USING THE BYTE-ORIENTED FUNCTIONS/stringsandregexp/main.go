package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	price := "€100"
	
	fmt.Println("Strings Prefix:", strings.HasPrefix(price, "€"))

	fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price), []byte{226, 130}))
}
