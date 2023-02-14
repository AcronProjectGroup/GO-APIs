package main

import "fmt"


func main() {


	name := "sina"
	space := " "
	sureName := "lalebakhsh"
	// combine two string: --------------<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
	var fullName strings.Builder
	fullName.WriteString(name)
	fullName.WriteString(space)
	fullName.WriteString(sureName)
	fmt.Println(fullName.String())

}
