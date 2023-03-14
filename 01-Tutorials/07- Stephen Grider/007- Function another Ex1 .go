// Source: Udemy tutorial from Stephen Grider

package main

import "fmt"

func main() {
    card := newCard()
    fmt.Println(card)

    card1, card2, card3 := newCard2()
    fmt.Println(card1, card2, card3)

}


func newCard() string {
    return "Five of Diamonds"
}


func newCard2() (firstString, secondString, thirdString string) {
    firstString = "First card"
    secondString = "Second card"
    thirdString = "Third card"
    return firstString, secondString, thirdString
}