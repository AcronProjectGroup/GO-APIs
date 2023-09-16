/*

	1- An unordered collection of key-value pairs

	2- map[Type of key]Type of value{}
	   map[string]float64{}

	3- provide a way to store data without relying on indexing

	4- This allows us to retrieve values based on their meaning and relation to other data types.

	5- maps have to be initialized before they can be used.

	6- it's not indexeble or ordered

	another names:
		1- associative array آرایه انجمنی
		2- hash table        جداول هش شده
		3- a dictionary      دیکشنری

*/

package main

import "fmt"

func main() {

	// var myNilMap map[string]string // nil map of string-int pairs

	// // LENGHT
	// fmt.Println(len(myNilMap))

	// // Why use nil map add can not add to it !!?

	// m1 := make(map[string]string) // Empty map of string-float64 pairs

	// m1["String"] = "3.14"

	// fmt.Printf("m1['String']= %q \n", m1["String"])

	// m2 := make(map[string]float64, 2) // Preallocate room for 100 entries

	// m2["strin1"] = 3.14
	// // m2["strin2"] = 5.21
	// // m2["strin3"] = 73.123

	// fmt.Println(m2)


	var myNilMap map[string]string // nil map of string-int pairs

	myRelativityMap :=  map[string]string{"string": "ssss"} // nil map of string-int pairs

	m3 := map[string]float64{       // Map literal
		"e":  2.71828,
		"pi": 3.1416,
		"pi1": 3.1416,
		"pi2": 3.1416,
		"pi3": 3.1416,
		"pi4": 3.1416,
		"pi5": 3.1416,
		"pi7": 3.1416,
		"pi8": 3.1416,
		"pi9": 3.1416,
		"pi10": 3.1416,
		"pi11": 3.1416,
		"pi12": 3.1416,
		"pi13": 3.1416,
		"pi14": 3.1416,
	}
	fmt.Println(len(m3))            // Size of map: 2

	m3["strin3"] = 5.12435
	fmt.Println(m3)


}
