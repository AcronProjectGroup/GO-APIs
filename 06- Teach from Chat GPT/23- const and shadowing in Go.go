// const and shadowing in Go

// Source Notion:  I don't had for this


package main

import (
	"fmt"
)

const Number int = 100
var PrintNumber string = fmt.Sprint(Number)

func main() {
	const Number int = 23
	fmt.Println(Number)
	fmt.Println(PrintNumber)
}

