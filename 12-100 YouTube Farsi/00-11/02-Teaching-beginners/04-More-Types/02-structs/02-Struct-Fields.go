package main

import "fmt"

type Boooks struct {
	Book1 string
	Book2 string
}

func main() {
	v := Boooks{"", "Qoran"}
	v.Book1 = "Choice of theory"
	fmt.Println(v.Book1)
	fmt.Println(v.Book2)

}
