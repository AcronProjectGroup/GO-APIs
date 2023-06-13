package main

import "fmt"

func printFibonacciSeries(counter_loop int) {
	numberOne := 0
	fmt.Printf("numberOne= %d\n", numberOne)

	numberTwo := 1
	fmt.Printf("numberTwo= %d\n", numberTwo)

	numberThree := numberTwo
	fmt.Printf("numberThree= %d\n", numberThree)

	//    fmt.Printf("Series is: %d %d", numberOne, numberTwo)

	counter_loop_plus := 1
	for true {
		fmt.Printf("------- loop number %d\n", counter_loop_plus)
		counter_loop_plus += 1
	
		numberThree = numberTwo
		fmt.Printf("numberThree = numberTwo => %d\n", numberThree)
		
		numberTwo = numberOne + numberTwo
		fmt.Printf("numberTwo = numberOne + numberTwo => %d\n", numberTwo)

		if numberTwo >= counter_loop {
			fmt.Println()
			break
		}
		numberOne = numberThree
		fmt.Printf("numberOne = numberThree %d\n", numberOne)


		// fmt.Printf(" %d", numberTwo)
	}
}

func main() {
	printFibonacciSeries(10)

}
