package main

import "fmt"

func main() {
   numbers := []int{2, 4, 6, 8, 10}

   fmt.Println("All numbers:", numbers)
   fmt.Println("First number:", numbers[0])
   fmt.Println("Last number:", numbers[len(numbers)-1])

   // تغییر مقدار یکی از عناصر برش
   numbers[3] = 12
   fmt.Println("Updated slice:", numbers)
}

// COPY :


package main

import "fmt"

func main() {
   numbers1 := [5]int{2, 4, 6}
   numbers2 := []int{8, 10}

   // Copying the elements of slice 2 to slice 1
   copy(numbers1[3:], numbers2)
   fmt.Println("Copied slice:", numbers1)

}