package main

import (
	"fmt"
	"bufio"
	"acron/animation"
	"strings"
	"os"
)


var sliceOfNames = []string{}

func main() {
	for {
		animation.Animation()
		var reader = bufio.NewReader(os.Stdin)
		fmt.Print("Write name -> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "finish" {
			break
		}
		sliceOfNames = append(sliceOfNames, userInput)
	}
	fmt.Println(sliceOfNames)
}

