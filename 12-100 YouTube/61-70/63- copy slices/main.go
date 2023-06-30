package main

import "fmt"

func main() {

	slice1 := []float64{}

	slice1 = []float64{3.14, 4.27, 7.31}
		
	slice10Cap := make([]float64, 10) //int = 0  string=""

	copy(slice10Cap, slice1)

	slice10Cap = append(slice10Cap, 44, 55, 66, 77)

	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice10 Capacity = ", slice10Cap)

}