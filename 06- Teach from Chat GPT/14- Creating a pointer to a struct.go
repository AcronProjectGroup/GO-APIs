// Creating a pointer to a struct

// Notion https://sinalalenakhsh.notion.site/7-Pointers-5bb883e6f3254a269fbc2105519a8193

package main

import "fmt"

type Person struct {
   Name string
   FamilyName string
   Age int
}

func main() {

   var Original_Struct Person = Person{"Sina", "LalehBakhsh", 30}
	//Many Codes -----
   var anotherPointer *Person = &Original_Struct

	var Person_Name string = (*anotherPointer).Name


   fmt.Println(Person_Name) 

   fmt.Println((*anotherPointer).FamilyName)

}

package main

import "fmt"

type Person struct {
   Name string
   Age int
}

func main() {
   var p Person = Person{"John", 30}
   var ptr *Person = &p
   fmt.Println((*ptr).Name) // Output: John
}