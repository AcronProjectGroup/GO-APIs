// https://quera.org/problemset/187845/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"strings"
)

func main() {


	var firstUserInput string
	var secondUserInput string

	userInput := bufio.NewReader(os.Stdin)
	finalUserInput, _ := userInput.ReadString('\n')
	finalUserInput = strings.TrimSuffix(finalUserInput, "\n")

	counter1 := 0
	counter2 := 0
	for _, charachters := range finalUserInput {
		trueUserNumber, err := strconv.Atoi(string(charachters))
		if err == nil && counter1 == 0 {
			firstUserInput += fmt.Sprint(trueUserNumber)
		} else if err != nil && counter1 == 0 {
			counter1 = 1
		} else if err == nil && counter2 == 0 {
			secondUserInput += fmt.Sprint(trueUserNumber)
		} else if err != nil && counter2 == 0 {
			counter2 = 1
			break
		} 
	}

	finalFirstUserInput , _:= strconv.Atoi(firstUserInput)
	finalSecondUserInput, _ := strconv.Atoi(secondUserInput)

	if finalFirstUserInput >= finalSecondUserInput && counter2 == 0{
		fmt.Println("Yes")
	} else if counter2 == 0 {
		fmt.Println("No")
	}

}




// #1 try
// fmt.Println(charachters)

// #2 try
// fmt.Println(string(charachters))

// #3 try
// fmt.Println(strconv.Atoi(string(charachters)))

/* Output
1 2 3
1 <nil>
0 strconv.Atoi: parsing " ": invalid syntax
2 <nil>
0 strconv.Atoi: parsing " ": invalid syntax
3 <nil>

*/
