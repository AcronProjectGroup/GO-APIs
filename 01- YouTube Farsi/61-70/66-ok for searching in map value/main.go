/*

* In many cases, this is undesirable and would lead to a bug in your program.
  When looking up the value in a map, Go can return a second, optional value.
  This second value is a bool and will be true if the key was found,
  or false if the key was not found.


* In Go, this is referred to as the ok idiom.
  Even though you could name the variable that captures
  the second argument anything you want,
  in Go, you always name it ok:

*/

package main

import (
	"fmt"
)

func main() {

	// Basic Search is not good:
		counts := map[string]int{ 
			"avalin kilid": 1,
			"2vomin kilid": 2,
		}

		counts1 := map[string]string{
			"ssss": "ssss",
		}
		fmt.Printf("%q\n", counts1["mmmm"])

		fmt.Println(counts["sammy"]) // = 0 Or if it is String value => "" <- stringe Khalii !!!
		fmt.Println(counts["avalin kilid"])
	

	// Best Practice
		count, ok := counts["sammy"]
		if ok {
			fmt.Printf("Sammy has a count of %d\n", count)
		} else {
			fmt.Println("Sammy was not found")
			// panic(count)
		}

}