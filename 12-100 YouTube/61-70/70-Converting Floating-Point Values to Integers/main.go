package main

import (
	"fmt"
	"math"
)

func main() {
	
	floatNumberToInteger := 19.50
	
	fmt.Println(math.Ceil(floatNumberToInteger))

	fmt.Println(math.Floor(floatNumberToInteger))
	
	fmt.Println(math.Round(floatNumberToInteger))
	
}
