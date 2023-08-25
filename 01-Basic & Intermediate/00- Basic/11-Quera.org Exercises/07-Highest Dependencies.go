package main

import "fmt"


func main()  {

	AllDuties := AllDuties()
	
	
}



func AllDuties() []int {
	var AllDuty []int 

	for i := 1; i <= 1000; i++ {
		AllDuty = append(AllDuty, i)
	}

	return AllDuty
}