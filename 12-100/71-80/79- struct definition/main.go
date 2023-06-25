package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Circle3Parameter struct {
	x float64
	y float64
	r float64
}

// Method to calculate the area of a Circle
func (c Circle) calculateArea() float64 {
	return math.Pi * c.radius * c.radius
}

// Function to calculate the area of a Circle
func calculateCircleArea(c Circle) float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	// Initialization: 
	// We can create an instance of our new Circle type in avariety of ways:
	var c Circle
	fmt.Println("c Circle: ",c)

	// We can also use the new function:
	circle2 := new(Circle)
	fmt.Println("circle2: ",circle2)

	// another initialization:
	circle3 := Circle3Parameter{x: 0, y: 0, r: 5}
	fmt.Println("circle3: ",circle3)

	// Fields
	// We can access fields using the . operator:
	fmt.Println("circle3: ", circle3.x, circle3.y, circle3.r)
	circle3.x = 10
	circle3.y = 5
	fmt.Println("circle3: ", circle3.x, circle3.y, circle3.r)



	// Creating a Circle instance using a struct
	circle := Circle{radius: 5}

	// Calculating the area using the struct method
	area1 := circle.calculateArea()
	fmt.Println("Area (using method):", area1)

	// Calculating the area using the function
	area2 := calculateCircleArea(circle)
	fmt.Println("Area (using function):", area2)
}