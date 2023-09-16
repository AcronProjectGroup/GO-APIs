/*

		I dont understand how it work ?  2023.Feb.22

*/



// // Source : https://go.dev/play/

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.


func main() {
	f := fib()
	fmt.Println(f())
	fmt.Println("-----")
	fmt.Println(f())
	fmt.Println("-----")
	fmt.Println(f())
	fmt.Println("-----")
	fmt.Println(f())
	fmt.Println("-----")

	for i := 0; i < 10; i++ {
	}
}

func fib() func() int {
	a, b := 0, 1
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return func() int {
		a, b = b, a+b
		fmt.Println("a=", a)
		fmt.Println("b=", b)
		return a
	}
}


