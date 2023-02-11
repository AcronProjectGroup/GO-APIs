package main

import (
	"fmt"
	"math"
)

func main() {
	var radius = 1.001
	area := math.Pi * radius * radius
	fmt.Println("area:", area)
}
