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
   
   // COPY :
   
	  numbers1 := [5]int{2, 4, 6}
	  numbers2 := []int{8, 10}
   
	  // Copying the elements of slice 2 to slice 1
	  copy(numbers1[3:], numbers2)
	  fmt.Println("Copied slice:", numbers1)



	  // مقایس آزایه ها با هم
	  // Arrays    
	arr1:= [3]int{9,7,6}
	arr2:= [...]int{9,7,6}
	arr3:= [3]int{9,5,3}
	fmt.Println(arr1==arr2)
	fmt.Println(arr2==arr3)
	fmt.Println(arr1==arr3)
}

