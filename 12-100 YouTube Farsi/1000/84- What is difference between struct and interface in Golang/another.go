/*

Interfaces:

    An interface is a collection of method signatures. It defines a set of methods that a type must implement to satisfy the interface.
    Interfaces provide a way to specify behavior without specifying the underlying concrete type.
    Any type that implements all the methods defined in an interface is said to satisfy that interface.
    Interfaces allow you to write more generic and flexible code by programming to an interface rather than a specific type.
    Multiple types can satisfy the same interface, allowing polymorphism.

Example of an interface:

*/

package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct {
	Name string
	DogBreed string // Rottweiler - German Shepherd - Shiba Inu
	Age int 
	
}

func (d Dog) Sound() string {
	return "Woof!"
}

type Cat struct {
	Name string
	CatBreed string // Sphynx cat
	Age int
}

func (c Cat) Sound() string {
	return "Meow!"
}

func main() {
	animals := []Animal{
		Dog{Name: "Buddy", DogBreed: "Rottweiler", Age: 2},
		Cat{Name: "Whiskers", CatBreed: "Sphynx cat", Age: 3},
	}

	for _, animal := range animals {
		fmt.Println(animal.Sound())
	}
}