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
	max := 10000

	// Step1:  Just get ENTER from User to first Step
	fmt.Println("")
	fmt.Println("If I guess your number ,write it and ENTER.")
	fmt.Println("")
	fmt.Printf("Select a number between 0 to %d, than press ENTER", max)
	var getEnter string = "\n"
	getENTER(getEnter)

	var counterGeuss int = 0
	for {
		// fmt.Println("min", min)
		// fmt.Println("max", max)
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(max-min) + min

		counterGeuss += 1
		fmt.Printf("\nCount my guess: %d\n\n", counterGeuss)
		fmt.Printf("I guess your number is %d.\n\nIf I wrong help me:\n\n", randomNumber)
		fmt.Printf("Your chosen number is greater than %d or less? bigger/smaller\n", randomNumber)

		var getUserAnswer string
		fmt.Scanln(&getUserAnswer)

		_, breakLoop := strconv.ParseInt(getUserAnswer, 0, 10)
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
			fmt.Printf("\nCount my guess: %d\n\n", counterGeuss)
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
