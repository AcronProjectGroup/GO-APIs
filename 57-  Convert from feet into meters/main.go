// 1 ft = 0.3048 m

package main

import "fmt"

func main() {

	/*
	var userInput_feet float64
	var result_feet_into_meters float64

	print("Enter number, convert from feet into meters: ")
	fmt.Scan(&userInput_feet)

	result_feet_into_meters = userInput_feet / 3.281
	fmt.Println("Result feet to meters is: ",result_feet_into_meters)
	*/

	var loop_number int
	print("how many times you want to calculate? ")
	fmt.Scan(&loop_number)
	for {
		if loop_number >= 5 {
			print("[!!] You cant choose more than 5\n")
			print("how many times you want to calculate? ")
			fmt.Scan(&loop_number)	
		} else {
			break
		} 
	}


	var user_Input float64
	print("What do you NUMBER want convert from feet into meters? ")
	fmt.Scan(&user_Input)

	result_feet_into_meters := user_Input / 3.281
	fmt.Printf("Result off %.f is this: %f meters.\n", user_Input, result_feet_into_meters)



}
















