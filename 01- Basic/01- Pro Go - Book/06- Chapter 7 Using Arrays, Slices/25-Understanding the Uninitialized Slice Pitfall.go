package main

import "fmt"

func main() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	
	var someNames []string
	
	copy(someNames, allNames)
	
	
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)

}

/*Output:

someNames: []
allNames [Lifejacket Paddle Hat]

*/



/*Because:

No elements have been copied to the destination slice. 
This happens because uninitialized slices have
zero length and zero capacity. 
The copy function stops copying when the length of the destination length
is reached, and since the length is zero, no copying occurs. 
No error is reported because the copy function
worked the way it is supposed to, but this is rarely the intended effect, 
and this is the likely cause if you
encounter an unexpectedly empty slice.

*/