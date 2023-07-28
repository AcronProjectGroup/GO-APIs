package main

import (
	"bufio"
	"os"
	"strings"
)

func GetReturn() string {
	UserInput := bufio.NewReader(os.Stdin)
	finalUserInput, _ := UserInput.ReadString('\n')
	finalUserInput = strings.TrimSuffix(finalUserInput, "\n")
	return finalUserInput
}





