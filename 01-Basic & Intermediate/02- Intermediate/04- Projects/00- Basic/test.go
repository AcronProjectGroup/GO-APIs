package main

import "fmt"

func Wierd() (namedOutput int) {
	defer func() {
		namedOutput = 5
	}()
	return 0
}

func main()  {
	wired := Wierd()
	fmt.Println(wired)
}