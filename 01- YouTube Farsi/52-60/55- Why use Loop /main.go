package main

import "fmt"


func main() {
	get_User_Input_Number := 1

	get_User_Input_anything := ""
	
	result := ""

	print("How many times do you want to play? ")
	fmt.Scan(&get_User_Input_Number)

	for {
		if get_User_Input_Number >= 10 {
			print("You can't choose more than 10\n")
			print("How many times do you want to play? ")
			fmt.Scan(&get_User_Input_Number)
	
		} else {
			break
		}
	}


	for i := 0; i < get_User_Input_Number; i++ {
		print(i+1,"-")
		fmt.Print("Tell me anything: ")
		fmt.Scan(&get_User_Input_anything)
		result = result + get_User_Input_anything + ".\n"
		get_User_Input_anything = ""
	}

	fmt.Print(result)

}


