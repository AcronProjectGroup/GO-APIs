package main

import "fmt"

var i, j int = 1, 2
func main() {
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
/*
output:
1 2 3 true false no!
*/ 


// BUT-----------------------------------
var i, j int = 1, 2
k := 3
func main() {
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
/*
output:
# example
./prog.go:6:1: syntax error: non-declaration statement outside function body
*/ 
