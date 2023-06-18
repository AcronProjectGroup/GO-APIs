/*

In many cases, this is undesirable and would lead to a bug in your program. 
When looking up the value in a map, Go can return a second, optional value. 
This second value is a bool and will be true if the key was found, 
or false if the key was not found. 

In Go, this is referred to as the ok idiom. 
Even though you could name the variable that captures 
the second argument anything you want, 
in Go, you always name it ok:

 */

package main

import "fmt"

func main() {

	counts := map[string]int{}
	fmt.Println(counts["sammy"])
	
	count, ok := counts["sammy"]
	if ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy was not found")
	}
}