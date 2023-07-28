package main

import (
	"fmt"
)


func main()  {
	
	finalUserInput := GetReturn()

	fmt.Println(finalUserInput)

	for index, value := range Sentences {
		fmt.Println(index,"----",value)
	}

}
