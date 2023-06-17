// 63- why use make function for create new variable of slices ?
// چرا از تابع میک برای ساختن آرایه استفاده میشود ؟


package main

import "fmt"

func main() {

	all_Scores := make([]float64, 3)
	
	just_Deriving_Scores := all_Scores[0:2]

	just_Deriving_Scores[0] = 12
	just_Deriving_Scores[1] = 13

	practical_Score := just_Deriving_Scores[0] // نمره عملی
	theory_Acore := just_Deriving_Scores[1] // نمره تئوری

	average_number_before_teacher := (practical_Score + theory_Acore) / float64(len(just_Deriving_Scores))


	if average_number_before_teacher >= 15 {
		fmt.Println("You won in driver testing :)")
	} else {
		fmt.Println("You lost in driver testing :(")
	}

	teacher_Grade_From_One_to_Five := all_Scores[2]

	if average_number_before_teacher < 15 {
		average_number_before_teacher += teacher_Grade_From_One_to_Five
		fmt.Println("Actually, your score was below 15,\nbut according to the teacher's opinion,\n5 points were added to you")
					// درواقع نمره شما زیر ۱۵ شده بود ولی با نظر معلم ۵ نمره به شما اضافه شد
		fmt.Println("\n\nSo You Won with Teacher Grade.\n\n")
	}



}

