package main

import (
	"fmt"
)

func main() {
	//             01234567
	courseName := "1A3B5X7H"

	for i := range courseName{
		fmt.Print(string(courseName[i]))
	}
}