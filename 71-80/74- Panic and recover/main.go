package main

import "fmt"

func main() {
	// panic("PANIC")
	// // the call to panic immediately stops execution of the function.
	// str := recover()
	// fmt.Println(str)

	defer func() {
		str := recover()
		fmt.Println(str)
		}()
		panic("PANIC")

}
