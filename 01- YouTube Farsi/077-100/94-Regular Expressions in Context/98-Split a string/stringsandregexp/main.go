package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "This  is  double  spaced"


	splitter := func(r rune) bool {
		return r == ' '
	}
	
	splits := strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}

	
}

/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "This  is  double  spaced"

	splits := strings.Fields(description)

	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}

}

*/

/* Output:

Field >>This<<
Field >>is<<
Field >>double<<
Field >>spaced<<

*/
