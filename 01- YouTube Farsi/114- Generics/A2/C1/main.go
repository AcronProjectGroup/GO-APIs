package main

import "fmt"

func AddIntegers(number1, number2 int) int {
	return number1 + number2
}

func AddFloats(number1, number2 float64) float64 {
	return number1 + number2
}


func AddAllNumbers[T int | float64](number1, number2 T) T {
	return number1 + number2
}


type Numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func AddAllInterface[T Numbers](number1, number2 T) T {
	return number1 + number2
}

func main()  {
	Result := AddIntegers(1, 2)
	fmt.Printf("%v\n", Result)

	resultSumFloats := AddFloats(1.1, 3.14)
	fmt.Printf("%v\n", resultSumFloats)

	resultAllNumbers := AddAllNumbers(1, 3.14)
	fmt.Printf("%v\n", resultAllNumbers)

	resultInterface := AddAllInterface(1, 1)
	fmt.Printf("%v\n", resultInterface)

}