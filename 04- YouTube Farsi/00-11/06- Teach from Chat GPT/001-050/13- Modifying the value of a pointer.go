// Modifying the value of a pointer

// Notion https://sinalalenakhsh.notion.site/7-Pointers-5bb883e6f3254a269fbc2105519a8193

package main

import "fmt"

func main() {
	var my_Original_Variable int = 234231

	// #####

	var myPointer *int = &my_Original_Variable

	*myPointer = 1

	fmt.Println(my_Original_Variable)
}
