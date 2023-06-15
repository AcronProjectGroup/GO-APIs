// Using the shorthand syntax to create a pointer to a struct

// Notion Source: https://sinalalenakhsh.notion.site/7-Pointers-5bb883e6f3254a269fbc2105519a8193

package main

import "fmt"

type Person struct {
   Name string
   Age int
}

func main() {
   var p Person = Person{"John", 30}

   var ptr = &p

   fmt.Println(ptr.Name) 
}
