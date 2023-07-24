package main

import "fmt"

type YekAdam struct {
	Name string
	SureName string
	Age int
}


func main()  {
	aghayeLaleh := YekAdam{
		// Ctrl + Space
		Name: "YourName",
		SureName: "Lalehbakhsh",
		Age: 30,
	}

	
	fmt.Println(aghayeLaleh.Name)
	fmt.Println(aghayeLaleh.SureName)
	fmt.Println(aghayeLaleh.Age)
	
	aghayeLALEH_Name := aghayeLaleh.Name
	fmt.Println(aghayeLALEH_Name)

	
}