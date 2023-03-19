package main

import (
	"fmt"
)

func main(){

	fmt.Println(add(1, 2))

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Today is not a weekday")
	}

	age := 10
	switch age {
	case 10:
		fmt.Println("Your age is 10.")
	case 11:
		fmt.Println("Your age is 11.")
	default:
		fmt.Println("Your age is not 10 or 11")
	}
	
}


func add(x, y int) int {
	return x + y
}

