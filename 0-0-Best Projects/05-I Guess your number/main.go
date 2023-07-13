package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	getBiggerOrSmaller := map[string]bool{
		"bigger":  true,
		"Bigger":  true,
		"BIGGER":  true,
		"smaller": false,
		"Smaller": false,
		"SMALLER": false,
	}
	min := 0
	max := 10

	// Step1:  Just get ENTER from User to first Step
	fmt.Println("")
	fmt.Println("If I guess your number ,write it and ENTER.")
	fmt.Println("")
	fmt.Printf("Select a number between 0 to %d, than press ENTER", max)
	var getEnter string = "\n"
	getENTER(getEnter)

	for {
		// fmt.Println("min", min)
		// fmt.Println("max", max)
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(max-min) + min

		fmt.Printf("I guess your number is %d.\n\nIf I wrong help me:\n\n", randomNumber)
		fmt.Printf("Your chosen number is greater than %d or less? bigger/smaller\n", randomNumber)

		var getUserAnswer string
		fmt.Scanln(&getUserAnswer)

		_, breakLoop := strconv.ParseInt(getUserAnswer, 0, 8)
		if breakLoop == nil {
			break
		}

		if getBiggerOrSmaller[getUserAnswer] == true {
			min = randomNumber + 1
		} else if getBiggerOrSmaller[getUserAnswer] == false {
			max = randomNumber
		}

		if min == max {
			fmt.Println("")
			fmt.Printf("I find your number!!! that is %d", min)
			fmt.Println("")
			break
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
