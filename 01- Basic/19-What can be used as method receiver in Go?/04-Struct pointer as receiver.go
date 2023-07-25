/*

Functions receivers can be also struct pointers. 
Remember that struct of given type 
and pointer to that struct are two different things in Go. 
As you may guess by passing a pointer to 
the struct as receiver type we give 
our method ability to modify all fields of the original struct.

*/

package main

import "fmt"

type A struct {
    i int
}

func (a *A) add() {
    a.i++
}

func (p *A) isAdult() bool {
    return p.i > 18
  }

func main() {
    a := A{i: 2}

    fmt.Printf("%#v\n", a)

    a.add()

    fmt.Printf("%#v\n", a)
}


/* Output:
main.A{i:2}
main.A{i:3}

Our function receives reference to original object and it is able to modify it. 
So you can see that value of field i gets modified from 2 to 3. Exactly what we expected.

*/

