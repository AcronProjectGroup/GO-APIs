package main

import "fmt"

func main()  {

	var myMultiDemensionalArray [3][3]int

	myMultiDemensionalArray[0][0] = 98

	myMultiDemensionalArray[1][1] = 456

	myMultiDemensionalArray[2][2] = 7777

	fmt.Println(myMultiDemensionalArray)
}

// Output
// [[98 0 0] [0 456 0] [0 0 7777]]


