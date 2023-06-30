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


	var myNilMap map[string]string // nil map of string-int pairs

	fmt.Println(myNilMap)

	// Why use nil map add can not add to it !!?

	// m1 := make(map[string]float64)      // Empty map of string-float64 pairs


	// m2 := make(map[string]float64, 100) // Preallocate room for 100 entries

	// m3 := map[string]float64{ // Map literal
	// 	"e":  2.71828,
	// 	"pi": 3.1416,
	// }
	// fmt.Println(len(m3)) // Size of map: 2

}
