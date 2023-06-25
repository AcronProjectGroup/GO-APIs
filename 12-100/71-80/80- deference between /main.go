/* 
The main difference between the two code snippets you provided lies in the function signature 
and how the Circle object is passed as an argument.
*/

// In the first code snippet:

/*

The circleArea function accepts a Circle object c as a value parameter. 
When you pass c to the function, 
a copy of the Circle object is made, 
and any modifications made to c within 
the function will not affect the original Circle object. 
Here, you're passing the Circle object by value.

*/

package main

import (
	"math"
	"fmt"
)

type Circle struct {
	x float64
	y float64
	r float64
}


func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}


func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))

	c2 := Circle{0, 0, 5}
	fmt.Println(circleArea(&c2))

}

/*

The circleArea function accepts a pointer to a Circle object c as a parameter. 
Instead of passing the entire Circle object, 
you're passing a memory address (pointer) to the Circle object. 
By using a pointer, 
any modifications made to c within the function will affect the original Circle object. 
Here, you're passing the Circle object by reference.

So, the main difference is in how the Circle object is passed to the function 
by value or by reference. 

If you pass it by value, a copy is made, 
and modifications won't affect the original object. 

If you pass it by reference (using a pointer), modifications will affect the original object.

*/