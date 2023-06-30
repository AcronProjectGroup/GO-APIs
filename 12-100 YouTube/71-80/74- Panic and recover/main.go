/*

A panic generally indicates a programmer error.



*/

package main

import "fmt"

func main() {
	// panic("PANIC")
	// //  But the call to recover will never happen in this case
	// //  because the call to panic immediately stops execution
	// //  of the function.
	// //  the call to panic immediately stops execution of the function.
	// str := recover()
	// fmt.Println(str)

	defer func() {
		str := recover()
		fmt.Printf("%s|\n",str)
		}()
		panic("PANIC")

}
