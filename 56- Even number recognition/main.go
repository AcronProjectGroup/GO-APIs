package main

import "fmt"


func main() {
	var userInput int64

	

	for i := 0; i < 3; i++ {
		
		fmt.Printf("loop number %d, Enter Your number: ", i+1)
		fmt.Scan(&userInput)

		remainder := userInput % 2
	
		if userInput % 2 == 0 {
			fmt.Printf("%d is even.\n", userInput)
		} else {
			fmt.Printf("%d is odd.\nand remainder of this number is %d.\n", userInput, remainder)
		}
	}

}






























// var number int

// number = 212312123 

// if number % 2 == 0 {
// 	println(number, "is even.")
// } else {
// 	println(number, "is odd.")
// }



