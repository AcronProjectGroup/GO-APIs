package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	x = math.Pow(999999999999999999999, 0.00000003)
	fmt.Println(x)
}
