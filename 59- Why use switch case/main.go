package main

import "fmt"

func main() {

	print("What is the name of your birthday month? just write number")
	print("1-farvardin\n2-ordibehesht\n3-khordad\n4-tir\n5-mordad\n6-shahrivar\n7-mehr\n8-aban\n9-azar\n10-dey\n11-bahman\n12-esfand\n")
	var userInput int
	fmt.Scan(&userInput)

	var rangeNumber map[int]bool

	rangeNumber = map[int]bool{
		1:  true,
		2:  true,
		3:  true,
		4:  true,
		5:  true,
		6:  true,
		7:  true,
		8:  true,
		9:  true,
		10: true,
		11: true,
		12: true,
	}

	for {
		if rangeNumber[userInput] == true {
			break
		} else {
			print("Enter range of number 1 to 12: ")
			fmt.Scan(&userInput)
		}
	}

}
