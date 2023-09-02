package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Print("")

	_, err := fmt.Scanf("%d %d", &num1, &num2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	Reminder := num1 % num2

	fmt.Println(Reminder)

}
