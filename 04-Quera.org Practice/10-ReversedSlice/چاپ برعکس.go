package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var AllNumbers []string	

	for {
		GetUserInput := bufio.NewReader(os.Stdin)
		FinalInput , err1 := GetUserInput.ReadString('\n')
		if err1 != nil {
			panic(err1)
		}
		FinalInput = strings.TrimSuffix(FinalInput ,"\n")
	
		if FinalInput == "0" {
			break
		}
		AllNumbers = append(AllNumbers, FinalInput )	
	}


	for i := range AllNumbers {
		fmt.Println(AllNumbers[len(AllNumbers)-1-i])
	}

}
