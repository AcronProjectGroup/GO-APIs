// Source: https://www.digitalocean.com/community/tutorials/defining-methods-in-go

package main

import "fmt"


func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	Creature.Greet(sammy)
}


type Creature struct {
	Name     string
	Greeting string
}

func (c Creature) Greet() {
	fmt.Printf("%s says %s", c.Name, c.Greeting)
}
