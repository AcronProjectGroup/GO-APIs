package main

import "fmt"

type A struct {
    i int
    b []byte
}

func (a A) modify() {
    a.i++
    a.b[0] = 10
}

func main() {
    a := A{b: make([]byte, 1)}

    fmt.Printf("%#v\n", a)
	fmt.Println(a.b)

    a.modify()

    fmt.Printf("%#v\n", a)
	fmt.Println(a.b)

}

/*

Output:
main.A{i:2, b:[]uint8{0x0}}
[0]
main.A{i:2, b:[]uint8{0xa}}
[10]



Note that when your struct contains reference as one of its members, 
then despite the fact that we pass a struct as copy the reference is the same. 
Another example illustrates it.

The output of programs turns out to be:
	main.A{i:2, b:[]uint8{0x0}}
	main.A{i:2, b:[]uint8{0xa}}


As you can see we passed struct to method as copy 
but method was able to modify original object pointer 
by reference inside a struct ([]byte slice 
in this example gets modified from {0x0} to {0xa})

*/