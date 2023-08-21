package main

import "fmt"

func main() {
	names := make([]string, 3, 6)

	names[0] = "Kayak"

	names[1] = "Lifejacket"

	// names[2] = "Paddle"

	fmt.Printf("%q \n", names)
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))



	myAppend := append(names, "one", "two", "three", "four", "five", "sechs")
	fmt.Printf("%q \n", myAppend)


	var anotherSlice []string

	anotherSlice = append(anotherSlice, "first")
	fmt.Printf("%q\n", anotherSlice)

}
