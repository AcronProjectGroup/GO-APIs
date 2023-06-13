package main

import "fmt"

func main() {
	
	print("What is the name of your birthday month? just write number")
	print("1-farvardin\n2-ordibehesht\n3-khordad\n4-tir\n5-mordad\n6-shahrivar\n7-mehr\n8-aban\n9-azar\n10-dey\n11-bahman\n12-esfand\n")
	var userInput int
	fmt.Scanln(&userInput)
	
	var rangeNumber map[int]int
	
	rangeNumber = map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
		9: 9,
		10: 10,
		11: 11,
		12: 12,
	}

	for i := 0; i <= 12-1 ; i++ {
		if rangeNumber[i] == userInput {
			break
		} else {

			print("You can just range of 1 to 12")
			print("1-farvardin\n2-ordibehesht\n3-khordad\n4-tir\n5-mordad\n6-shahrivar\n7-mehr\n8-aban\n9-azar\n10-dey\n11-bahman\n12-esfand\n")
			var userInput int
			fmt.Scan(&userInput)
						
		}

	}

}