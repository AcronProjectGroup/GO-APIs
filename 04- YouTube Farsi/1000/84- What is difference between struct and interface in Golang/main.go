/*In Go, both structs and interfaces are important concepts that serve different purposes. Here's an explanation of the differences between structs and interfaces:

Structs:

    A struct is a composite data type that allows you to group together zero or more values with different types into a single entity.
    It represents a collection of fields, where each field has a name and a type. It defines the structure and layout of data.
    Structs are used to create custom data types that can have their own properties and behavior.
    Structs are value types, meaning when you assign a struct variable to another variable or pass it as a function argument, a copy of the struct is created.
    Structs can have methods associated with them, allowing you to define behaviors and actions that can be performed on struct instances.

Example of a struct: */

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "John Doe",
		Age:  25,
	}

	fmt.Println(person.Name) // Output: John Doe
	fmt.Println(person.Age)  // Output: 25
}