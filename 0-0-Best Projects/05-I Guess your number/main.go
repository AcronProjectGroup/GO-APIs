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
	
	// Step1:  Just get ENTER from User to first Step
	var getEnter string = "\n"
	fmt.Printf("Select a number between 0 to %.f, than press ENTER", endNumber)
	getENTER(getEnter)
	
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(int(endNumber) - int(startNumber) + 1)
		var getUserAnswer string
		fmt.Printf("more than %d?\n", randNum)
		fmt.Scanln(&getUserAnswer)

		if getYesOrNo[getUserAnswer] == true {
			startNumber = endNumber / 2
			startNumber += 1
			fmt.Println("startNumber", startNumber)
			fmt.Println("endNumber", endNumber)
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
