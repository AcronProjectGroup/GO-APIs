package main

import "fmt"

func main()  {
	var getUserInput float64
	fmt.Scanf("%f", &getUserInput)
	
	for {
		result := int(getUserInput) % 2 

		if result == 0 {
			getUserInput = getUserInput / 2
			fmt.Print(getUserInput, "-")
			if getUserInput == 1 {
				break
			}
		} else {
			getUserInput = (3 * getUserInput) + 1  
			fmt.Print(getUserInput, "-")
			if getUserInput == 1 {
				break
			}
		}
	}

	fmt.Println()
}