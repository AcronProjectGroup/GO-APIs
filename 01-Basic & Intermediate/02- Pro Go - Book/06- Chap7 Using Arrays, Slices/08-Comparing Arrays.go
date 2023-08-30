// The comparison operators == and != can be applied to arrays

package main

import "fmt"

func main() {
	
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	
	moreNames := [3]string{"Kayak", "Lifejacket", "Paddle"}
	
	same := names == moreNames
	
	fmt.Println("comparison:", same)
	
}



// Output:
// comparison: true