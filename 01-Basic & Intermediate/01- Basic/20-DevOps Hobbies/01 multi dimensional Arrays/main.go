package main

import "fmt"

func main()  {


	var my3DArray [2][3][4]int64
	
	count := 0

	for index1 := 0; index1 < 2; index1++ {
		for index2 := 0; index2 < 3; index2++ {
			for index3 := 0; index3 < 4; index3++ {
				count += 1
				my3DArray[index1][index2][index3] = int64(index1) + int64(index2) + int64(index3)				
				fmt.Println(count,"---",my3DArray)
			}
		}
	}



}