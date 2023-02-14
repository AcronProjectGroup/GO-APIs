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
	for i := 0; i <= 2; i++ {
		animation.Animation()
		var reader = bufio.NewReader(os.Stdin)
		fmt.Print("Write name -> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		sliceOfNames = append(sliceOfNames, userInput)
	}
	fmt.Println(sliceOfNames)
}

func getForAddToSlice(a string)  {
	
}