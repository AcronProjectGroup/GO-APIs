// Passing a pointer to a function

//  Notion Source: https://sinalalenakhsh.notion.site/7-Pointers-5bb883e6f3254a269fbc2105519a8193

// ارسال اشاره گر به یک تابع:



package main

import "fmt"

func addOne(p *int) {
   *p += 1
}

func main() {
   var num int = 42
   addOne(&num)
   fmt.Println(num) // Output: 43
}
// --------------------------
// --------------------------
// --------------------------
// --------------------------


package main

import "fmt"

func addOne(myPointer *int) {
   *myPointer += 1
}

func main() {
   var myNumber int

   fmt.Println("Enter a number:")
   fmt.Scan(&myNumber)

   addOne(&myNumber)

   fmt.Println("The number plus one is:", myNumber)
}
