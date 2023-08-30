package main

import (
	"fmt"
	"unsafe"
)

type IntBool struct {
	a int64
	b bool
	c int32
}

type BoolInt struct {
	a bool
	b int64
	c int32
}

func main()  {
	if unsafe.Sizeof(IntBool{}) == unsafe.Sizeof(BoolInt{}) {
		fmt.Println("Same")
	} else {
		fmt.Println("Different")
	}
	
}


/*


func Sizeof ¶

func Sizeof(x ArbitraryType) uintptr

Sizeof takes an expression x of any type and returns the size in bytes of a hypothetical variable v as if v was declared via var v = x. 

The size does not include any memory possibly referenced by x. 

For instance, if x is a slice, Sizeof returns the size of the slice descriptor, 
not the size of the memory referenced by the slice. 

For a struct, the size includes any padding introduced by field alignment. 

The return value of Sizeof is a Go constant if the type of the argument x does not have variable size. 

(A type has variable size if it is a type parameter or if it is an array or struct type with elements of variable size).


*/