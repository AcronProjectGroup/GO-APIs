/*
The cars are produced on an assembly line. 
The assembly line has a certain speed, that can be changed. 
The faster the assembly line speed is, the more cars are produced. 
However, changing the speed of the assembly line also changes 
the number of cars that are produced successfully, 
that is cars without any errors in their production.

Implement a function that takes in the number of cars produced per hour 
and the success rate and calculates the number of successful cars made per hour. 
The success rate is given as a percentage, from 0 to 100:

CalculateWorkingCarsPerHour(1547, 90)
// => 1392.3

Note: the return value should be a float64.

*/

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	last := float64(productionRate) * successRate / 100
    return last
}