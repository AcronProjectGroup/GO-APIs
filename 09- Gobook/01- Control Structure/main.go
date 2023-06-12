package main


import "fmt"

func main() {
	userInputNumber := 0
	userInputAnything := ""

	print("How many times do you want to play? ")
	fmt.Scan(&userInputNumber)

	print("Tell me anything: ")
	fmt.Scan(&userInputAnything)


	fmt.Printf("You want %d times to play.\nAnd you told me [%s].\n", userInputNumber, userInputAnything)

}
