/*

Array:
	1- index start from 0
	2- lenght start from 1
	3- ordered
	4- capacity define when created

Slice:
	1- index start from 0
	2- lenght start from 1
	3- ordered
	4- There is no limit on the capacity and there is no need to define the capacity
    	محدودیت در ظرفیت نداره و لزوما نیازی به تعریف ظرفیت نداره
		یعنی اینکه میشه براش ظرفیت تعیین کرد
*/


package main

import "fmt"

func main() {

	// Basic Create
		var myArray [2]int
		fmt.Println("myArray: ", myArray)
		myArray[0] = 1
		myArray[1] = 1
		fmt.Println("myArray: ", myArray)

	// Basic Create
		var mySlice []int    
		fmt.Println("mySlice: ", mySlice) 
		mySlice = append(mySlice, 1)
		mySlice = append(mySlice, 2)
		mySlice = append(mySlice, 3)
		fmt.Println("mySlice: ", mySlice) 

	// Slice with fixed capacity: 
		// A slice of a certain length without populating -  use the built-in make() function:
		mySlice_withـfixedـcapacity :=  make([]int, 3)
		fmt.Println("Slice with fixed capacity: ",mySlice_withـfixedـcapacity)
		mySlice_withـfixedـcapacity = append(mySlice_withـfixedـcapacity, 1, 2, 3)
		fmt.Println("Slice with fixed capacity: ",mySlice_withـfixedـcapacity)

		mySlice_withـfixedـcapacity[0] = 1
		mySlice_withـfixedـcapacity[1] = 2
		mySlice_withـfixedـcapacity[2] = 3
		fmt.Println("Slice with fixed capacity: ",mySlice_withـfixedـcapacity)

}