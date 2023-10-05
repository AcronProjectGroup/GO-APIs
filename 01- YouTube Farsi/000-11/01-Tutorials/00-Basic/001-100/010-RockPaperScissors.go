package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func main() {

	playerChoice := PlayerChoice() 
	fmt.Println("Your Choice:", playerChoice)

	computerChoice := ComputerChoice()
	fmt.Println("Computer Choice:", computerChoice)

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

func ComputerChoice() string{
	//create the computerChoiceSlice slice and append computerChoice to it
	computerChoiceSlice := make([]string, 0)
	computerChoiceSlice = append(computerChoiceSlice,
    "rock",
    "paper",
    "paper",)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	computerChoice := fmt.Sprint(computerChoiceSlice[rand.Intn(len(computerChoiceSlice))])
	return computerChoice
}