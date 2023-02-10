package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {

	playerChoice := PlayerChoice() 

	
	fmt.Println("Your Choice:", playerChoice)
}

func PlayerChoice() string{
	for {
		var reader = bufio.NewReader(os.Stdin)
		fmt.Print("Please enter rock, paper, or scissors -> ")
		playerChoice, _ := reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)
		if playerChoice == "rock" || playerChoice == "paper" || playerChoice == "scissors" {
			return playerChoice
		} 	
	}
}