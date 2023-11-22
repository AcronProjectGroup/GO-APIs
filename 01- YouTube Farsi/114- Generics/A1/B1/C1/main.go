package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
	// current directory Terminal Or CMD
	// go mod tidy
)

func AddEveryTypes[T constraints.Ordered](number1, number2 T) T {
	return number1 + number2
}

func main()  {
	ResultAddEverything := AddEveryTypes(1, 3.1412334)
	fmt.Printf("%v\n", ResultAddEverything)
}