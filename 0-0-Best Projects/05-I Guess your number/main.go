package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var endNumber float64 
var startNumber float64

func main() {
	var endNumber float64 = 1001

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
	randNum := rand.Intn(int(endNumber))

	var correctGuess bool = false
	
	getYesOrNo := make(map[string]bool)
	getYesOrNo["yes"] = true
	getYesOrNo["Yes"] = true
	getYesOrNo["YES"] = true
	getYesOrNo["no"] = false
	getYesOrNo["No"] = false
	getYesOrNo["NO"] = false
	
	var getUserAnswer string
	
	for correctGuess != true {
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

	if getYesOrNo[getUserAnswer] == true {
		startNumber = endNumber / 2
		rand.Seed(time.Now().UnixNano())
		startNumber += 1	
		rand.Intn(int(endNumber) - int(startNumber) +1)
	}


}









