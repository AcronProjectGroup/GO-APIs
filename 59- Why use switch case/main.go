package main

import (
	"fmt"
)

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

	if userInput == 1 {
		fmt.Println("Happy birthday! you are Farvarnini... :))")
	} else if userInput == 2 {
		fmt.Println("Happy birthday! you are Ordibeheshti... :))")
	} else if userInput == 3 {
		fmt.Println("Happy birthday! you are Khordadi... :))")
	} else if userInput == 4 {
		fmt.Println("Happy birthday! you are Tirii!!... :))")
	} else if userInput == 5 {
		fmt.Println("Happy birthday! you are Mordadi... :))")
	} else if userInput == 6 {
		fmt.Println("Happy birthday! you are Shahrivari... :))")
	} else if userInput == 7 {
		fmt.Println("Happy birthday! you are Mehri... :))")
	} else if userInput == 8 {
		fmt.Println("Happy birthday! you are Abani... :))")
	} else if userInput == 9 {
		fmt.Println("Happy birthday! you are Azari... :))")
	} else if userInput == 10 {
		fmt.Println("Happy birthday! you are Deyii!... :))")
	} else if userInput == 11 {
		fmt.Println("Happy birthday! you are Bahmani... :))")
	} else if userInput == 12 {
		fmt.Println("Happy birthday! you are Esfandi... :))")
	}

	switch userInput {
	case 1:
		fmt.Println("Happy birthday! you are Farvarnini... :))")
	case 2:
		fmt.Println("Happy birthday! you are Ordibeheshti... :))")
	case 3:
		fmt.Println("Happy birthday! you are Khordadi... :))")
	case 4:
		fmt.Println("Happy birthday! you are Tirii!!... :))")
	case 5:
		fmt.Println("Happy birthday! you are Mordadi... :))")
	case 6:
		fmt.Println("Happy birthday! you are Shahrivari... :))")
	case 7:
		fmt.Println("Happy birthday! you are Mehri... :))")
	case 8:
		fmt.Println("Happy birthday! you are Abani... :))")
	case 9:
		fmt.Println("Happy birthday! you are Azari... :))")
	case 10:
		fmt.Println("Happy birthday! you are Deyii!... :))")
	case 11:
		fmt.Println("Happy birthday! you are Bahmani... :))")
	case 12:
		fmt.Println("Happy birthday! you are Esfandi... :))")
	}

}
