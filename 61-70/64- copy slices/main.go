package main

import "fmt"

func main() {

	twoNumbers := []float64{1,2}

	thirdNumbers := make([]float64, 3)

	copy(thirdNumbers, twoNumbers)

	fmt.Println(thirdNumbers)

	thirdNumbers[2] = 3

	fmt.Println(thirdNumbers)



}