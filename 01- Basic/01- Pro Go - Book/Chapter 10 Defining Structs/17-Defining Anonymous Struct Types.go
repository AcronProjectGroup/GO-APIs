// Anonymous struct types are defined without using a name

package main

import "fmt"


// struct is parameter of function writeName
func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
}



func main() {
	
	type Product struct {
		name, category string
		price          float64
	}
	
	type Item struct {
		name     string
		category string
		price    float64
	}
	
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Stadium", category: "Soccer", price: 75000}
	
	writeName(prod)
	writeName(item)
}
