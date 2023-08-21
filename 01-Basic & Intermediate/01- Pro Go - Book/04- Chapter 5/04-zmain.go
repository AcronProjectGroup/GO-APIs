package main

import (
	"fmt"
	"math"
)

func main() {
	positiveResult := 3 % 2

	negativeResult := -3 % 2

	absoluteResult := math.Abs(float64(negativeResult))

	fmt.Println("positiveResult:", positiveResult)

	fmt.Println("negativeResult:", negativeResult)

	fmt.Println("absoluteResult:",absoluteResult)
}
