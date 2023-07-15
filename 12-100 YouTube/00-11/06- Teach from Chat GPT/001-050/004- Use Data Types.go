package main

import (
	"fmt"
)


func main(){
	x := 10
	y := 5

	sum := x + y // arithmetic operator
	greater := x > y // comparison operator
	result := (x > y) && (x < 20) // logical operator

	fmt.Println(sum)
	fmt.Println(greater)
	fmt.Println(result)

	age := 36

	if age >= 18 && age <= 35 {
		fmt.Println("You are in 18 to 35.")
	} else if age <= 0 {
		fmt.Println("You are not born!!!")
	} else if age > 0 && age < 18 {
		fmt.Println("You are a young.")
	} else {
		fmt.Println("You are a old.")
	}


}