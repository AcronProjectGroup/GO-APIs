package main

import "fmt"

type mostatil struct {
	tool float64 
	arz float64
}

func (Mostatil mostatil) area() float64 {
	return Mostatil.arz * Mostatil.tool
}

type dayereh struct {
	shoAa float64
}

func (Dayereh dayereh) area() float64 {
	return 3.14 * Dayereh.shoAa * Dayereh.shoAa 
}

type myInterface interface {
	area() float64 
}

func PrintMyInterface(MyIN myInterface) {
	fmt.Printf("Area is %.1f\n", MyIN.area())
}

func main()  {
	MosMan := mostatil{tool: 2, arz: 3}

	DayMan := dayereh{shoAa: 10}

	PrintMyInterface(MosMan)

	PrintMyInterface(DayMan)



}



