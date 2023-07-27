package main

import "fmt"


type AllInOne interface {
	returnCountryName() string
	returnCityName() string
	getCountryArea() float64
}


func main()  {
	allInOne := []AllInOne{
		City{"Tehran", "Iran"},
		Area{"Washington.D.C","USA",177.0},
	}


	for _, oneOfAll := range allInOne {
		fmt.Println(
			"City:", oneOfAll.returnCityName(),
			"Country:", oneOfAll.returnCountryName(), 
			"Area:", oneOfAll.getCountryArea(),"km",
		)
	}

}