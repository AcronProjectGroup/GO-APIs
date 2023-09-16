package main

import "fmt"

func main() {

	// یه مدل دیگه تعریف کردن اسلایس
	threeNumbers := []float64{}

	// به چه دردی میخوره این مدل تعریف کردن؟؟
	threeNumbers = []float64{1.56, 2.34, 5.872}


	// why use make() ? because get "Minimum capacity"
	tenNumbers := make([]float64, 10)

	// copy()
	copy(tenNumbers, threeNumbers)

	fmt.Println("three Numbers: ", threeNumbers)
	fmt.Println("ten Numbers:   ", tenNumbers)

	// What if ... ? from slice into array ?
	var twoNumbers [2]float64 
	copy(twoNumbers[:], threeNumbers)
	fmt.Println("two Numbers:   ", twoNumbers)

	// What if ... ? from slice into another slice ?
	twoNumbers_Slice := make([]float64, 2) 
	copy(twoNumbers_Slice, threeNumbers)
	fmt.Println("two Numbers Slice:   ", twoNumbers_Slice)



	// Result:
		// append => add another elements
		// copy   => change exist elements

}