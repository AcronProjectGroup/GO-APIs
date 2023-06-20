/*

 1- if our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before both of them
 2- deferred functions are run even if a run-time panic occurs.
 3-


*/

package main

import (
	"fmt"
	"os"
)


func first() {
	fmt.Println("1st Avali")
}
func second() {
	fmt.Println("2nd Dovomi")
}
func main() {
	defer first()
	second()

	f, _ := os.Open("filename.txt")
	defer f.Close()
}
