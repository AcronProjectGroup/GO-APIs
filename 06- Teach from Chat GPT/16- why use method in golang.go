// why use method in golang ? 

// Notion Source: https://sinalalenakhsh.notion.site/8-why-use-method-in-golang-1e880de696b84938836a160ecf599902

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (person Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.", person.Name, person.Age)
}

func (person Person) Goodbye() {
	fmt.Printf("Goodbye, my name is %s and I am %d years old.\n", person.Name, person.Age)
}

func main() {
	myPerson1 := Person{"Sina", 30}
	myPerson1.Greet()
}
