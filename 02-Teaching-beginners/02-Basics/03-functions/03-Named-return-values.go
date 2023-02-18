package main

import "fmt"

func split(sum int) (asghar, y int) {
	asghar = sum * 4 / 9
	y = sum - asghar
	return
}

func main() {
	fmt.Println(split(17))
}
