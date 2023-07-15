package main


// Functon more example


import "fmt"

func main() {
    card := newCard()
    fmt.Println(card)

    
    card1, card2 := newCard2()
    fmt.Println(card1, card2)


    card1, card2, card3 := newCard3()
    fmt.Println(card1, card2, card3 )

}


func newCard() string {
    return "Just one return string"
}

func newCard2() (firstString, secondString string) {
    firstString = "First-card"
    secondString = "Second-card"
    return firstString, secondString
}

func newCard3() (firstString, secondString, thirdString string) {
    firstString = "First-card"
    secondString = "Second-card"
    thirdString = "Third-card"
    return firstString, secondString, thirdString
}
