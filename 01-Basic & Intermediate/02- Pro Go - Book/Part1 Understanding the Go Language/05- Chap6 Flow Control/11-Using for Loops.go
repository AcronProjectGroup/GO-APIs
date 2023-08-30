package main

import (
	"fmt"
	//"strconv"
)

func main() {
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break
		}
	}
}
