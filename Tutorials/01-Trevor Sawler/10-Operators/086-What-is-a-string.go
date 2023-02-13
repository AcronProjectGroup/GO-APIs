package main

import "fmt"


func main() {
	name := "sina"
	for i := 0; i < len(name); i++ {
		fmt.Printf("i %x \n", name[i])
	}
}
