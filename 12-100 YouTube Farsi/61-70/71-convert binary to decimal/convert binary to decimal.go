package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

func main() {

	finalUserInput := getUserInput()
	// finalUserInput = 1010

	// reverse string
	finalUserInput = reverseString(finalUserInput)
	// Result 	  = 0101

	// Base2  	  = 2222
	var base2 float64 = 2

	// Result 	  = 0101
	// Base2  	  = 2222
	// Position   = 1234
	var resultConvertToDecimal int
	for position, number := range finalUserInput {
		integerNumber, _ := strconv.Atoi(string(number))
		resultConvertToDecimal += int(integerNumber) * int(math.Pow(base2, float64(position)))
	}

	fmt.Println(resultConvertToDecimal)
}

func getUserInput() string {
	userInput1 := bufio.NewReader(os.Stdin)
	finalUserInput, _ := userInput1.ReadString('\n') // 1010101000101010\n
	finalUserInput = strings.TrimSuffix(finalUserInput, "\n")

	return finalUserInput
}


// function to reverse string
func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _ , v := range str {
		result = string(v) + result

	}
	return
}
