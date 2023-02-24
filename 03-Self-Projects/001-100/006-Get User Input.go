// Source: https://www.tutorialspoint.com/how-to-get-input-from-the-user-in-golang

package main

import (
	"fmt"
	// "strings"
)


func main() {
	var getUserInput string

	fmt.Println(&getUserInput)

	fmt.Println("Start with 'lower','upper','title'")
	fmt.Println("Write your opinion:")

	fmt.Scan(&getUserInput)

	if string(getUserInput[0:5]) == "lower" {
		fmt.Println("First user = lower")
	}
	

}



