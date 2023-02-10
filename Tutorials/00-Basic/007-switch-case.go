package main

import (
	"fmt"

)


func main() {
	one := 1
	two := 2
	three := 3
	switch three + one {
	case one + one:
		fmt.Println("one + one:", one + one)
	case two + two:
		fmt.Println("two + two:", two + two)
	case two + one:
		fmt.Println("two + one", two + one)
	case three + three:
		fmt.Println("three + three", three + three)
	}
}