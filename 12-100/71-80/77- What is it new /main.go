package main

import "fmt"



func theFirstMode(number1 int) {
	number1 = 0
}

func theSecondMode(number2 *int) {
	*number2 = 6
}

func theThirdMode(number3 *int) {
	*number3 = 1
}

func main() {
	number1 := 5
	theFirstMode(number1)
	fmt.Println("number1: ",number1) // number1 is still 5
	fmt.Println("&number1: ",&number1) // memory location 
	// fmt.Println("*number1: ",*number1) // can't be use

	// -----------------
	number2 := 5
	theSecondMode(&number2) // need ues & --> without & we get error
	fmt.Println("number2: ",number2) // number2 changes to 6
	fmt.Println("&number2: ",&number2) // memory location
	// fmt.Println(*number2) // Can't be use 

	// ----------------- 
	number3 := new(int)
	// theThirdMode(&number3) // don't need use &
	theThirdMode(number3)
	fmt.Println("number3: ",number3)  
	fmt.Println("&number3: ",&number3)  // memory location
	fmt.Println("*number3: ",*number3) // ?

	number4 := number3
	println("number4: ",number4)

	number5:= *number3
	println("number5: ",number5)

}



