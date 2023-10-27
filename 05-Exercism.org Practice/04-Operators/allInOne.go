package main

import "fmt"

func main() {
	// last := CalculateWorkingCarsPerHour(1371, 90)
	// fmt.Println(last)

	// another := CalculateWorkingCarsPerMinute(1105, 90)
	// fmt.Println(another)

	sina := CalculateCost(37)
	fmt.Println(sina)
}

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	last := float64(productionRate) * successRate / 100
	return last
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	last := (float64(productionRate) * successRate / 100) / 60
	return int(last)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10
	individualCars := carsCount % 10
	costIs := (groupsOfTen * 95000) + (individualCars * 10000)
	return uint(costIs)
}
