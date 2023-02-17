package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 12
	b := 5

	check_result, error_ := divideTwoNumbers(a, b)
	// In this logical if first condition passed b != 0 than second condition is checked.
	if error_ != nil {
		fmt.Println(error_)
	} else if b % 2 == 0 {
		fmt.Println(a, "/", b, "=", check_result )
		fmt.Println("So", a, "is even number.")
		
	} else {
		fmt.Println(a, "/", b, "=", check_result )
		fmt.Println(b, "can't divide exactly", a)
		
	}

	

}

func divideTwoNumbers(original_number, divider_number int) (int, error) {
	if divider_number == 0 {
		return 0, errors.New("divider number can't be zero")
	} 
	return original_number / divider_number, nil
}