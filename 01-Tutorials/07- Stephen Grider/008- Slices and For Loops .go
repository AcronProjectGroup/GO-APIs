package main

import "fmt"

func main() {

	arrayes := [4]string{"avali", "dovomi", "sevomi", "chaarami"}
	fmt.Println(arrayes)

	cardsAce := []string{}

	cards := []string{"Ace Pik!", card()}

	// Adding to slices
	cardsAce = append(cards, "Ace Khesht!!!")

	fmt.Println(cards)

	fmt.Println(cardsAce)
}

func card() string {
	return "Ace Dell!!"
}

/* Output:

[Ace Pik! Ace Dell!!]
[Ace Pik! Ace Dell!! Ace Khesht!!!]

*/
