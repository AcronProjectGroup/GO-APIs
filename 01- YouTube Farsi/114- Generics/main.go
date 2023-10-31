package main

import "fmt"

type CustomMap[T comparable, V int | string] map[T]V

func main() {

	obj := make(CustomMap[int, string])
	obj[3] = "3"
	obj[1] = "1"

	fmt.Println(obj)
}