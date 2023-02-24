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

	fmt.Println(string(getUserInput[0]))
	fmt.Println(getUserInput[0])

}



