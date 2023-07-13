package main

import (
	"fmt"
	"math"
)

func main() {
	// kayak := 275
	
	floatNumberToInteger := 19.50
	fmt.Println(math.Ceil(floatNumberToInteger))
	floatNumberToInteger = 19.5
	fmt.Println(math.Floor(floatNumberToInteger))
	floatNumberToInteger = 19.5
	fmt.Println(math.Round(floatNumberToInteger))


	// total := kayak + int(math.Round(soccerBall))
	
	// fmt.Println(total)
}

/*

Ceil(value)		This function returns the smallest integer that is greater than the specified floating-
				point value. The smallest integer that is greater than 27.1, for example, is 28.

Floor(value)	This function returns the largest integer that is less than the specified floating-point
				value. The largest integer that is less than 27.1, for example, is 28.

Round(value)	This function rounds the specified floating-point value to the nearest integer.


RoundToEven(value) 	This function rounds the specified floating-point value to the nearest even integer.


*/