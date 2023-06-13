// 1 ft = 0.3048 m

package main

import "fmt"

func main() {

	var userInput_feet float64
	var result_feet_into_meters float64

	print("Enter number, convert from feet into meters: ")
	fmt.Scan(&userInput_feet)

	result_feet_into_meters = userInput_feet / 3.281
	fmt.Println("Result feet to meters is: ",result_feet_into_meters)
}