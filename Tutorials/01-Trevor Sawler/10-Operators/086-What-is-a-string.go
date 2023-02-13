package main

import "fmt"


func main() {
	name := "sina"
	for i := 0; i < len(name); i++ {
		fmt.Printf("i %x \n", name[i])
	}

	fmt.Println("----------------------")
	fmt.Println("Index\tRune\tString")
	fmt.Println("-----\t----\t------")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	name = "δικο τους"
	fmt.Println("----------------------")
	fmt.Println("Index\tRune\tString")
	fmt.Println("-----\t----\t------")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

}
