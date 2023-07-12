package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// var test123 float64 = 1001
	// rand.Seed(time.Now().UnixNano())
	// randnum := rand.Intn(int(test123))
	// fmt.Println(randnum)

	var rangeNumber float64 = 1001
	var getEnter string = "\n"
	fmt.Println("Select a number between 0 to 1000, than press ENTER")

	for {
		fmt.Scanln(&getEnter)
		if getEnter != "\n" {
			getEnter = "\n"
			fmt.Println("Just press ENTER!!!")
		} else {
			break
		}
	}

	// Set the seed value for the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate a random Int type number between 0 and 9
	randNum := rand.Intn(int(rangeNumber))

	var correctGuess bool = false

	for correctGuess != true {
		getYesOrNo := make(map[string]bool)
		getYesOrNo["yes"] = true
		getYesOrNo["Yes"] = true
		getYesOrNo["YES"] = true
		getYesOrNo["no"] = false
		getYesOrNo["No"] = false
		getYesOrNo["NO"] = false

		var getUserAnswer string
		fmt.Println("Your guess is great than", randNum/2, "?")
		fmt.Scanln(&getUserAnswer)
		getUserAnswer = strings.TrimSpace(getUserAnswer)

		if getUserAnswer != "\n" {
			var getTrueAnswer bool = false
			for getTrueAnswer == false {
				for index := range getYesOrNo {
					if getUserAnswer == index {
						getTrueAnswer = true
						break
					}
				}
				if getTrueAnswer == true {
					break
				} else {
					fmt.Println("Your guess is great than", randNum/2, "?")
					fmt.Scanln(&getUserAnswer)
					getUserAnswer = strings.TrimSpace(getUserAnswer)
				}

			}
		}

		correctGuess = true

	}

}
