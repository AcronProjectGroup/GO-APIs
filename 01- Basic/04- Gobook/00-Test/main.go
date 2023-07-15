package main

import "fmt"

func theFirstMode(numberOne int) {
	numberOne = 0
}

func theSecondMode(numberTwo *int) {
	*numberTwo = 0
}

func theThirdMode(numberThree *int) {
	*numberThree = 1
}

func main() {
	numberOne := 5
	theFirstMode(numberOne)
	fmt.Println(numberOne) // numberOne is still 5

	// -----------------
	numberTwo := 5
	theSecondMode(&numberTwo)
	fmt.Println(numberTwo) // x is 0

	// ----------------- 
	// numberThree := new(int)
	// theThirdMode(numberThree)
	// fmt.Println(*numberThree) // numberThree is 1
	// fmt.Println(&numberThree) // 

}

