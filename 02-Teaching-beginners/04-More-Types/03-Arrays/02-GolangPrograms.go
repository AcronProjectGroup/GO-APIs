// Source : https://www.golangprograms.com/go-language/golang-maps.html



package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intArray [5]int
	var strArray [5]string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(intArray))

	fmt.Println(reflect.ValueOf(strArray).Kind())
	fmt.Println(strArray)

}


