package main

import "fmt"



func main()  {

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
		if getUserAnswer != "\n" {
			for index := range getYesOrNo {
				fmt.Println(index)
			}
		}

		correctGuess = true
	}

	

}











