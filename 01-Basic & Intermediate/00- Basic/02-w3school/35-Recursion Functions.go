package main

import (
	"fmt"
)

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func main() {
	testcount(1)
}

// Result:
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
