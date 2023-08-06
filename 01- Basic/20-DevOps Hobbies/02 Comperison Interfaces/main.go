package main

import "fmt"

func main()  {
	myInteger := 100
	comparisonTypeOfInterface(myInteger)

	var mySlice = []string{"sina",}
	comparisonTypeOfInterface(mySlice)

	myMap := map[string]int{
		"Sina":1,
		"Mahsaan":1,
	}
	comparisonTypeOfInterface(myMap)
}



func comparisonTypeOfInterface(myInterface interface{}){
	switch myInterface.(type) {
	case bool:
		fmt.Printf("%v is Boolean.\n",myInterface)
	case string:
		fmt.Printf("%v is String.\n",myInterface)
	case float64:
		fmt.Printf("%v is Float64.\n",myInterface)
	case float32:
		fmt.Printf("%v is Float32.\n",myInterface)
	case int:
		fmt.Printf("%v is Integer.\n",myInterface)
	case rune:
		fmt.Printf("%v is Rune.\n",myInterface)
	case interface{}:
		fmt.Printf("%v is Interface.\n",myInterface)
	default:
		fmt.Println("Unknown type.")
	}

}