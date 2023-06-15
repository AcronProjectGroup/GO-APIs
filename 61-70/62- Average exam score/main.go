// One of the reasons for using slices is the number of fixed indexes
//  یکی از دلایل استفاده از اسلایس تعداد ایندکس ثابت است

package main

import "fmt"

func main() {
	var driving_Exam [2]float64

	println("Enter the practical driving test score (1 to 20): ")
	fmt.Scan(&driving_Exam[0])

	println("Enter the driving theory test score (1 to 20): ")
	fmt.Scan(&driving_Exam[1])

	var total float64 = 0
	// for i := 0; i < len(driving_Exam); i++ {
	// 	total += driving_Exam[i]
	// }

	for /* Or i */ _ , value := range driving_Exam {
		// if i > len(driving_Exam) {
		// 	println(i)
		// }
		total += value
	}


	if total / float64(len(driving_Exam)) > 15 {
		Average_Exam_Score := total / float64(len(driving_Exam))
		fmt.Printf("Your average exam score driving is %.1f.\nYou passed the exam.\n", Average_Exam_Score)
	} else if  total / float64(len(driving_Exam)) <= 15 {
		Average_Exam_Score := total / float64(len(driving_Exam))
		fmt.Printf("Your average exam score driving is %.1f.\nYou did not pass the exam.\n", Average_Exam_Score)
	}

}
