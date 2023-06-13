package main


func main() {
	// var number int

	// number = 212312123 

	// if number % 2 == 0 {
	// 	println(number, "is even.")
	// } else {
	// 	println(number, "is odd.")
	// }
	fibunaci()
}




func fibunaci() {
	/*
	1
	1+1=2
	    2+1=3
			3+2=5
				5+3=8
					8+5=13
	*/
								// 
	first := 0  				// 
	second := 1 				// 1 
	third := second 			// 1
	for i := 0; i < 6; i++ {
		third = second 			// 1
		second = first + second // 1
		first = third 			// 1
		println(i, "- thir", third)
	}
	println(" ")

}