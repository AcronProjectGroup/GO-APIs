// Source: https://www.golangprograms.com/golang/methods-interfaces-objects/


package main

import "fmt"

func main() {
	num := 2
	mul:= multiply(num)
	
	multiToNumber := multiToNumber(2)
	
	
	
	fmt.Println("Ten times of a given number is: ",mul.tentimes())
	fmt.Println("multiToNumber:",multiToNumber)
}

func multiToNumber(number int) int {
	return number * 2
}


type multiply int	

func (m multiply) tentimes() int {
	return int(m * 10)
}


