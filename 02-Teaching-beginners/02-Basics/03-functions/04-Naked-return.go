package main

import "fmt"

func split(sum int) (asghar, y int) {
	asghar = sum * 4 / 9
	y = sum - asghar
	return
}

func main() {
	asghar, y := split(3)


	fmt.Println("Asghar:", asghar, "y:", y)
}



/*

this returns behavior :
returning any variables difined in function for outputing.

*/ 