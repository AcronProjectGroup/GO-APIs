package main

import "fmt"

func main() {
	first := 100
	
	second := &first
	
	first++
	
	fmt.Println("First:", first)
	
	fmt.Println("Second:", *second)


	var myFirst int

	var mySecond = &myFirst

	fmt.Println(*mySecond)


}


