package main

import (
	"fmt"
	"os"
)

func main() {
	// Example for variables
	var variable string
	const constant1 = "string"
	fmt.Println(variable)
	fmt.Println(constant1)
	var x string = "Hello World"
	fmt.Println(x)

	// Example for numbers
	fmt.Println("1 + 1 =", 1+1)

	// Example for boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Defining Multiple Variables
	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)

	// Get UserInput
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println(input)

	// for loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// if
	if i%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// switch and case and default
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

	// array
	var x123 [5]int
	fmt.Println(x123)

	// another array
	var x333 [5]int
	x333[4] = 100
	fmt.Println(x333)

	// range and value
	var xfloat [5]float64
	var total float64 = 0
	for _, value := range xfloat {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	// slice
	var xSlice []float64
	fmt.Println(xSlice)

	// make function
	xMake := make([]float64, 5)
	fmt.Println(xMake)

	// 27.
	// Slice Functions
	// append
	xMake = append(xMake, 3.14)
	xMake = append(xMake, 3.14)
	xMake = append(xMake, 3.14)
	fmt.Println(xMake)
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// 28.
	// [low : high]
	arr := [5]float64{1, 2, 3, 4, 5}
	xAnother := arr[0:5]
	fmt.Println(xAnother)

	// 29.
	// map
	var xMap map[string]int
	fmt.Println(xMap)
	// If you were to read this out loud you would say “xis a map of strings to ints.”

	// 30.
	// slices maps
	var xSlicesMaps map[string]int
	xSlicesMaps["key"] = 10
	fmt.Println(xSlicesMaps["key"])

	// 31.
	// delete function
	delete(xSlicesMaps, "key")

	// 32.
	// make map
	xMakeMap := make(map[int]int)
	xMakeMap[1] = 10
	fmt.Println(xMakeMap[1])

	// 34.
	// Your Second Function
	var sliceForIt []float64
	average(sliceForIt)

	// 35.
	// Returning Multiple Values
	f()

	// 36.
	// Variadic Functions
	add(1, 1, 1, 1, 1, 1)

	// 37.
	// Closure
	add37 := func(x, y int) int {
		return x + y
	}
	fmt.Println(add37(1, 1))

	// 38.
	// increment
	xInc := 0
	increment := func() int {
		xInc++
		return xInc
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// 39.
	// return function from a function
	makeEvenGenerator()

	// 40.
	// Recursion
	factorial(4)

	// 42.
	defer second()
	first()
	/* Description
	This program prints 1st followed by 2nd. Basically de-fer moves the call to second to the end of the function:
	func main() {
		first()
		second()
	}
	*/

	// 43.
	f, _ := os.Open("fileName.txt")
	defer f.Close()

	// 45.

	/* False:
	func main() {
		panic("PANIC")
		str := recover()
		fmt.Println(str)
	}
	*/

	// True:
	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	// panic("PANIC")

	// func zero(x int) {
	// 	x = 0
	// }
	xDontChanges := 5
	zero(xDontChanges)
	fmt.Println(xDontChanges) // x is still 5









	
}






















// 34.
func average(xs []float64) float64 {
	panic("Not Implemented")
}

// 35.
func f() (int, int) {
	return 5, 6
}

// 36.
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// 39.
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// 40.
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

/*
• Is x == 0? No. (x is 2)
• Find the factorial of x – 1
	• Is x == 0? No. (x is 1)
	• Find the factorial of x – 1
		• Is x == 0? Yes, return 1.
	• return 1 * 1
• return 2 * 1
*/

// 42.
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func zero(x int) {
	x = 0
}
