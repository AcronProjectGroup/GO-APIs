/*

Array:
	1- index start from 0
	2- lenght start from 1
	3- ordered
	4- capacity define when created
	5- Limit by one type



/* Slice:
	1- index start from 0
	2- lenght start from 1
	3- ordered
	4- لزوما نیازی به تعیین ظرفیت در زمان تعریف ندارند
	5- Limit by one type
*/



package main

import "fmt"

func main() {
	// Basic Create
		var myArray [2]int // [lenght]type [capacity]type
		// var myStringArray [2]string
		fmt.Println("myArray: ", myArray)
		// fmt.Printf("myArray: %q\n", myStringArray)
		myArray[0] = 11
		myArray[1] = 22
		fmt.Println("myArray: ", myArray)
	
	// Basic Create
		var mySlice []int    
		fmt.Println("mySlice: ", mySlice) 
		mySlice = append(mySlice, 1) // [ 1, 13, 12, 14, 2341, 13 , append is here]
		// mySlice = append(mySlice, 2)
		// mySlice = append(mySlice, 3)
		mySlice[0] = 10
		fmt.Println("mySlice: ", mySlice) 

	// Slice with reserved elements:   برش با عناصر رزرو شده
		// use the built-in make() function:
		mySlice_withـfixedـcapacity :=  make([]int, 3)
		fmt.Println("Slice with fixed capacity: ",mySlice_withـfixedـcapacity)
		mySlice_withـfixedـcapacity = append(mySlice_withـfixedـcapacity, 1, 2, 3)
		fmt.Println("Slice with fixed capacity: ",mySlice_withـfixedـcapacity)
		// چجوری اون سه تای اول رو تغییر بدیم؟
		// How can I access to they?
		mySlice_withـfixedـcapacity[0] = 1
		mySlice_withـfixedـcapacity[1] = 2
		mySlice_withـfixedـcapacity[2] = 3
		fmt.Println("Slice with fixed capacity: ",mySlice_withـfixedـcapacity)

	// Slicing Arrays into Slices
		// note however, that when slicing an array, the result is a slice, not an array.
		mySlice_from_myArray := myArray[0:]
		fmt.Printf("Type(mySlice from myArray): %T\n", mySlice_from_myArray)
		mySlice_from_myArray = append(mySlice_from_myArray, 6,7,8,9)
		fmt.Println("mySlice from myArray: ", mySlice_from_myArray)
	
}