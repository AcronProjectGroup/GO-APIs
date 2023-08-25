package main

import "fmt"


func main()  {

	GetNumberOfSoups()
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