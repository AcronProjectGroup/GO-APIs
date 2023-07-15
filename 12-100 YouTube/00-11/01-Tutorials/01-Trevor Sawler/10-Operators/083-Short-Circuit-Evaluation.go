package main

import "fmt"

func main() {
	a := 12
	b := 6


	// In this logical if first condition passed b != 0 than second condition is checked.
	if b != 0 && divideTwoNumbers(a, b) == 2 {
		fmt.Println(a, "/", b, "=", divideTwoNumbers(a, b) )
		fmt.Println("So", a, "is even number.")
	}

}

func divideTwoNumbers(s, m int) int {
	return s / m
}