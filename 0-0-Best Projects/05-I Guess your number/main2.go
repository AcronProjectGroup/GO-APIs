package main

import (
	"fmt"
	"math/rand"
	"time"

)



func main()  {

	rangeNumber := 6

	nesf123 := rangeNumber / 2
	rand.Seed(time.Now().UnixNano())
	nesf123 += 1
    fmt.Println(rand.Intn(rangeNumber - nesf123 +1) + nesf123)

	var getEnter string = "\n"
	fmt.Println("Select a number between 0 to 1000, than press ENTER")

	// Set the seed value for the random number generator
	rand.Seed(time.Now().UnixNano())
	// Generate a random Int type number between 0 and 9
	randNum := rand.Intn(int(endNumber))

	var correctGuess bool = false

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
					if getYesOrNo[getUserAnswer] == true {
						startNumber = endNumber / 2
						rand.Seed(time.Now().UnixNano())
						startNumber += 1
					}
				}

			}
		}

		// correctGuess = true
	}

}











