package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main()  {

	Number, err1 := GetNumberOfSoups()
	if err1 != nil {
		fmt.Println(err1)
	}

	SliceStrings, err2 := SendSliceOfString(Number)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(SliceStrings)

}



func GetNumberOfSoups() (int, error) {
	var NumberUserInput int
	_, err := fmt.Scanf("%d", &NumberUserInput)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	} else if NumberUserInput <= 1 || NumberUserInput >= 100001 {
		fmt.Println("Error:", err)
		return 0, err

	}
	return NumberUserInput, nil
}


func SendSliceOfString(Number int) ([]string, error) {
	var SliceStrings []string
	for i := 1; i <= Number; i++ {
		userInput := bufio.NewReader(os.Stdin)
		finalUserInput, _ := userInput.ReadString('\n')
		finalUserInput = strings.TrimSuffix(finalUserInput, "\n")
		SliceStrings = append(SliceStrings, finalUserInput)
	}
	return SliceStrings, nil
}
