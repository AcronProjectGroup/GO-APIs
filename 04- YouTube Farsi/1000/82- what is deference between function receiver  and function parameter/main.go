/*
	The difference between a function receiver and a function parameter in Go lies in their roles and usage within a function.

	Function Receiver:

		* A function receiver is associated with a method, which is a function defined on a specific type (struct, interface, etc.).
		* It is defined within the method signature and placed between the func keyword and the method name.
		* The receiver specifies on which type the method operates and allows the method to be called on instances of that type.
		* The receiver is similar to the concept of this or self in other programming languages.
		* It provides access to the fields and methods of the type on which the method is defined.
		* The receiver can be of two types: value receiver (func (t T) methodName()) or pointer receiver (func (t *T) methodName()), depending on whether the method modifies the receiver or not.

	Function Parameter:

		* A function parameter is a variable defined within the parentheses of a function declaration.
		* It is used to receive values that are passed to the function when it is called.
		* Parameters define the input values that a function expects to receive and work with during its execution.
		* They are defined with a name and a type, and multiple parameters can be listed, separated by commas.
		* Parameters are local variables within the function's scope and can be accessed and used within the function body.

	Here's an example to illustrate the difference between a function receiver and a function parameter:
*/

package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

// Method with receiver
func (r Rectangle) calculateArea() float64 {
	return r.width * r.height
}

// Function with parameters
func calculateRectangleArea(width, height float64) float64 {
	return width * height
}

func main() {
	rect := Rectangle{width: 4, height: 3}

	// Calling method on the Rectangle type using the receiver
	area1 := rect.calculateArea()
	fmt.Println("Area calculated using method:", area1)

	// Calling function with parameters
	area2 := calculateRectangleArea(rect.width, rect.height)
	fmt.Println("Area calculated using function:", area2)
}


