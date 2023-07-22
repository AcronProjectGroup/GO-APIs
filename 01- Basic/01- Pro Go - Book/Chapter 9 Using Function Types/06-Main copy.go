package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var countryCapital = map[string]string{
	"USA":     "Washington, D.C",
	"Germany": "Berlin",
}

func justGetAndReturn(fromUserInput string) string {
	return fromUserInput
}

func getCapital_OK(justGetAndReturn string, country map[string]string) (capital string, ok bool) {
	capital, ok = country[justGetAndReturn]
	return capital, ok
}


func resultGetCapital_OK()(capital string, nil){
	return capital
}

func main() {

	fmt.Println("USA or Germany? ")
	userInput := bufio.NewReader(os.Stdin)
	finalUserInput, _ := userInput.ReadString('\n')
	finalUserInput = strings.TrimSuffix(finalUserInput, "\n")

	if capital, ok := getCapital_OK(finalUserInput, countryCapital); ok {
		fmt.Printf("Capital of %s's %v.\n", finalUserInput, capital)
	}

}
