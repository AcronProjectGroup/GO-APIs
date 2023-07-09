package main

import (
	"fmt"
)

const (
	first = iota + 1
	second
	_
	fourth
)

func main() {
	fmt.Println(first, second, fourth)
}
