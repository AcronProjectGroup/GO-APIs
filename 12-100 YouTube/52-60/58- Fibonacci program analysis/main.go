// 58- Fibonacci program analysis
// تحلیل عدد فیبوناچی که به زبان گو نوشته شده



package main

import "fmt"

func printFibonacciSeries(counter_loop int) {
	numberOne := 0
	// fmt.Printf("numberOne= %d\n", numberOne)

	numberTwo := 1
	// fmt.Printf("numberTwo= %d\n", numberTwo)
	fmt.Printf("Fibinacci Series= %d", numberTwo)

	numberThree := numberTwo
	// fmt.Printf("numberThree= %d\n", numberThree)

	//    fmt.Printf("Series is: %d %d", numberOne, numberTwo)

	// counter_loop_plus := 1
	for i := 1 ; i <= counter_loop-1 ; i++ {
		// fmt.Printf("------- loop number %d\n", counter_loop_plus)
		// counter_loop_plus += 1
	
		numberThree = numberTwo
		// fmt.Printf("numberThree = numberTwo => %d\n", numberThree)
		
		numberTwo = numberOne + numberTwo
		// fmt.Printf("numberTwo = numberOne + numberTwo => %d\n", numberTwo)
		fmt.Printf(" , %d ", numberTwo)

		// if numberTwo >= counter_loop {
		// 	fmt.Println()
		// 	break
		// }
		numberOne = numberThree
		// fmt.Printf("numberOne = numberThree => %d\n", numberOne)


		// fmt.Printf(" %d", numberTwo)
	}
}

func main() {

	var counter_loop int
	println("How many Fibonacci series do you want to have? ")
	fmt.Scan(&counter_loop)

	printFibonacciSeries(counter_loop)

}
