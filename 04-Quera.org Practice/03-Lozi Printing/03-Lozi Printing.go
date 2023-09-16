package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Print("")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if n%2 == 0 {
		fmt.Println("Please enter an odd number.")
		return
	}

	// Upper half of the pattern
	for i := 0; i < n/2; i++ {
		spaces1 := strings.Repeat(" ", (n/2)-i)
		stars := strings.Repeat("*", 2*i+1)
		spaces2 := strings.Repeat(" ", 2*(n/2-i))

		fmt.Println(spaces1 + stars + spaces2 + stars + spaces1)
	}

	// Middle line
	stars := strings.Repeat("*", n)
	spaces := strings.Repeat("", n/2)
	fmt.Println(stars + spaces + stars)

	// Lower half of the pattern
	for i := n/2 - 1; i >= 0; i-- {
		spaces1 := strings.Repeat(" ", (n/2)-i)
		stars := strings.Repeat("*", 2*i+1)
		spaces2 := strings.Repeat(" ", 2*(n/2-i))

		fmt.Println(spaces1 + stars + spaces2 + stars + spaces1)
	}
}



