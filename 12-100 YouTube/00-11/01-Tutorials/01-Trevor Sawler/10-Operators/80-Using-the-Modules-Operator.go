package main

import "fmt"

func main() {
	for month := 1; month <= 12; month ++ {
		fmt.Println("The month after", month, "is", month % 12 + 1)
	}
}
