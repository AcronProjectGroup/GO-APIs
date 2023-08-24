package main

import (
	"fmt"
)

func main() {


	// TotalUserInput, Newerr1 := GetFirst()
	// if Newerr1 != nil {
	// 	fmt.Printf(Newerr1.Error())
	// }

	// SliceNumber, Newerr2 := GetUserAddInSlice(TotalUserInput)
	// if Newerr2 != nil {
	// 	fmt.Printf(Newerr2.Error())
	// }

	// GetSliceNumbers(SliceNumber)

}

func GetFirst() (int, error) {
	var num1 int
	fmt.Print("")
	_, err := fmt.Scanf("%d", &num1)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, err
	} else if num1 < 1 || num1 >= 20001 {
		fmt.Println("Error:", err)
		return 0, err

	}

	TotalUserInput := num1 * 4

	return TotalUserInput, nil
}

func GetUserAddInSlice(TotalUserInput int) ([]int, error) {
	var SliceNumber []int
	for i := 0; i < TotalUserInput; i++ {
		var num2 int
		_, err := fmt.Scanf("%d", &num2)
		if err != nil {
			fmt.Println("Error:", err)
			return SliceNumber, err
		}
		SliceNumber = append(SliceNumber, num2)
	}
	return SliceNumber, nil
}

func GetSliceNumbers(originalSlice []int) {
	var LastSlice []int
	sliceSize := 4
	for len(originalSlice) >= sliceSize {
		currentSlice := originalSlice[:sliceSize]

		first := currentSlice[0]
		second := currentSlice[1]
		// third := currentSlice[2]
		// fourth := currentSlice[3]

		convertor := 0
		if second < first {
			convertor = second // 1
			second = first     // 2
			first = convertor  // 1
		}

		// 	// 1 1 1 1
		// if first == second && first == third && first == fourth {
		// 	LastSlice = append(LastSlice, 0)
		// 	// 2 1 1 2
		// }

		// if first > second && second == third && third < fourth {
		// 	LastSlice = append(LastSlice, 0)
		// 	// 2 1 3 2
		// }

		// if first > second && second < third && third > fourth {
		// 	LastSlice = append(LastSlice, 1)
		// 	// 1 1 1 4
		// }

		// if first == second && second == third && third < fourth {
		// 	LastSlice = append(LastSlice, 1)
		// 	// 1 4 1 1
		// }

		// if first < second && second > third && third == fourth {
		// 	LastSlice = append(LastSlice, 1)
		// }

		originalSlice = originalSlice[sliceSize:]
	}

	for SINA := range LastSlice {
		fmt.Println(LastSlice[SINA])
	}

}
