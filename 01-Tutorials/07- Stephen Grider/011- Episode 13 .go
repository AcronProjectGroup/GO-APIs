package main

import "fmt"

func main() {
	originalNumbers := []int{2, 3, 5, 11, 13, 1, 12, 43, 65, 666, 192,}
	       //                0  1  2  3   4   5  6   7   8   9    10

	// COPY :
	numbers_Of_Original_Numbers := []int{0, 1, 2, 3}
	copy(originalNumbers[3:], numbers_Of_Original_Numbers[0:3])
	
	fmt.Println(originalNumbers)
	fmt.Println(numbers_Of_Original_Numbers)

	numbers := []int{2, 3, 5, 11, 13, 1, 12, 43, 65, 666, 192,}

	// Change a one Index in Slices
	fmt.Println(numbers)
	numbers[0] = 6266
	fmt.Println(numbers)

	// InvalidLen occurs when an argument to the len built-in function is not of
	// supported type.
	fmt.Println(len(numbers[3]))

	// Reverse Indexing
	fmt.Println("Last number:", numbers[len(numbers)-2])

	// Length 1, 2, 3, 4,  5,  6
	fmt.Println(len(numbers))

	// Indexing
	fmt.Println(numbers[2])
}
