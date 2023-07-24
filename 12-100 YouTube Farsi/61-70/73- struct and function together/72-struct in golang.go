package main

import "fmt"

type YekAdam struct {
	Name string
	SureName string
	Age int
}

func meghdarDehiyeYekAdam (Name, SureName string, Age int) *YekAdam {
	return &YekAdam{Name, SureName, Age}
}



func main()  {
	aghayeLaleh := meghdarDehiyeYekAdam("SINA", "LALE", 30)

	fmt.Println(aghayeLaleh.Name)
	fmt.Println(aghayeLaleh.SureName)
	fmt.Println(aghayeLaleh.Age)

	
}