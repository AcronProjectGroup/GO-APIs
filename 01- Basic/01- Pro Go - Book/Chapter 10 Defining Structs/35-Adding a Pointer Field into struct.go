package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	*Supplier
}


type Supplier struct {
	name, city string
}



func main() {
	var prod Product
	var prodPtr *Product
	fmt.Println("Value:", prod.name, prod.category, prod.price, prod.Supplier.name)
	fmt.Println("Pointer:", prodPtr)
}

/*

The problem here is the attempt to access the name field of the embedded struct. 
The zero value for the embedded field is nil, which causes the following runtime error:

	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xc0000005 code=0x0 addr=0x0 pc=0x5bc592]
	goroutine 1 [running]:
	main.main()
	        C:/structs/main.go:20 +0x92
	exit status 2


*/