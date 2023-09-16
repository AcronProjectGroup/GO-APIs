package main

import (
	"fmt"
)

func main() {
	str := "I like Go 😀"
	var counter int

	for i := 0; i < len(str); i++ {
		counter++
	}
	fmt.Println(counter)
	fmt.Println([]byte("😀"))
	fmt.Println([]byte(str))


}
