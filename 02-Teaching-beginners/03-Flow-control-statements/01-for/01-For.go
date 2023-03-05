package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)
}

/*
Output:
0
1
3
6
10
15
21
28
36
45
45

*/ 