/*

	1- An unordered collection of key-value pairs

	2- map[Type of key]Type of value{}
	   map[string]float64{}

	3- provide a way to store data without relying on indexing

	4- This allows us to retrieve values based on their meaning and relation to other data types.
	   این به ما امکان می دهد مقادیر را بر اساس معنی و ارتباط آنها با سایر انواع داده ها بازیابی کنیم.

	5- maps have to be initialized before they can be used. 

	6- it's not indexeble or ordered  

	another names:
		1- associative array
		2- hash table
		3- adictionary

*/

package main

import "fmt"


func main() {
	// Basic Create:
		// var x map[string]int
		// x["key"] = 10
		// fmt.Println(x["key"])
		
		xPlus := make(map[string]int)
		xPlus["key"] = 10
		fmt.Println(xPlus["key"])
 
	// delete an elements in map and slice :
		var mySlice []int
		myMap := make(map[string]string) 

		mySlice = append(mySlice, 1,2,3,4)
		
		myMap["yek"] = "yek"
		myMap["do"] = "do"
		myMap["se"] = "se"

		fmt.Println(mySlice)
		fmt.Println(myMap)
		fmt.Printf("%q\n", myMap)

		delete_element_mySlice := append(mySlice[:2], mySlice[3:]...)
		fmt.Println(delete_element_mySlice)

		delete(myMap, "do")
		fmt.Println(myMap)

}