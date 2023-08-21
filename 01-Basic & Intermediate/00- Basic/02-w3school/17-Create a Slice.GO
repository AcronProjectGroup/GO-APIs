/*


Syntax
slice_name := []datatype{values}

A common way of declaring a slice is like this:
myslice := []int{}

*/

package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}

// Result:
// 0
// 0
// []
// 4
// 4
// [Go Slices Are Powerful]
