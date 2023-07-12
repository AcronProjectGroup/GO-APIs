package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var endNumber float64 = 10
	var startNumber float64 = 0
	getYesOrNo := map[string]bool{
		"yes": true,
		"Yes": true,
		"YES": true,
		"no":  false,
		"No":  false,
		"NO":  false,
	}
	randNum := rand.Intn(int(endNumber) - int(startNumber) + 1)
	if randNum == 0 {
		randNum = rand.Intn(int(endNumber) - int(startNumber) + 1)
	}

	// Step1:  Just get ENTER from User to first Step
	var getEnter string = "\n"
	fmt.Printf("Select a number between 0 to %.f, than press ENTER", endNumber)
	getENTER(getEnter)

	for i := 0; i < 10; i++ {
		var getUserAnswer string
		fmt.Printf("more than %d?\n", randNum)
		fmt.Scanln(&getUserAnswer)

		if getYesOrNo[getUserAnswer] == true {
			// startNumber = endNumber / 2
			// startNumber += 1
			randNum = rand.Intn(int(endNumber) - int(startNumber) )
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
