package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	getBiggerOrSmaller := map[string]bool{
		"bigger": true,
		"Bigger": true,
		"BIGGER": true,
		"smaller":  false,
		"Smaller":  false,
		"SMALLER":  false,
	}
	min := 0
	max := 10

	// Step1:  Just get ENTER from User to first Step
	var getEnter string = "\n"
	fmt.Printf("Select a number between 0 to %d, than press ENTER", max)
	getENTER(getEnter)



	for i := 0; i < 10; i++ {
		fmt.Println("min", min)
		fmt.Println("max", max)
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(max - min + 1) + min
		fmt.Printf("Your chosen number is greater than %d or less? bigger/smaller\n", randomNumber)

		var getUserAnswer string
		fmt.Scanln(&getUserAnswer)

		if getBiggerOrSmaller[getUserAnswer] == true {
			min = randomNumber + 1
		} else if getBiggerOrSmaller[getUserAnswer] == false {
			max = randomNumber
		}
	}

}

func getENTER(getEnter string) {
	getEnter = "\n"
	for {
		fmt.Scanln(&getEnter)
		if getEnter != "\n" {
			getEnter = "\n"
			fmt.Println("Just press ENTER!!!")
		} else {
			break
		}
	}

}
