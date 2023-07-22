package main

import "fmt"

func main() {
	
             
	for number := 0; number <= 3; number ++ {
		fmt.Println(number)
	} 

	for initialization; condition; postcondition {
        // do something
	}


	//--------------------------------------------------------------

	originalNumbers := []string{"2", "341", "5", "11", "13", "1", "12", "43"}
	//                           0    1      2    3     4     5     6     7   

	for index_number, index_Value := range originalNumbers {
		fmt.Println(index_number, "=", index_Value + ".000")
	}

	originalNumbers123 := []int{2, 3, 5, 11, 13, 1, 12, 43, 65, 666, 192,}
	       //                0  1  2  3   4   5  6   7   8   9    10

	for i := range originalNumbers123 {
		fmt.Println(i)
	}


	for index, whatIsIt  := range originalNumbers123 {
		fmt.Println(index, ":", whatIsIt)
	}



}
