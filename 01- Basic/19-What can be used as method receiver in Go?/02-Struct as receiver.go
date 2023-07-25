// Struct as receiver

package main

import "fmt"

type A struct {
	i int
}

//  Notice that these methods receive struct copy as their receiver argument. 
func (a A) add() {
	a.i++
}


//  Notice that these methods receive struct copy as their receiver argument. 
func (a A) get() int {
	return a.i
}

func main() {
	a := A{i: 2}
	a.add()
	fmt.Println(a.get())
}

/* Output:
2

As you may expect the output of this program is 2. 
This is because add() method modifies the copy of our struct not our original struct.
So the conclusion, of when to utilize struct as receiver type, 
is to use it when you don’t want to allow any modifications of your original object.


*/
