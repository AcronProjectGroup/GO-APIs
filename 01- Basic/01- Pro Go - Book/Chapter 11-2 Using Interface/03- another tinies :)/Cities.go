package main

type City struct {
	CityName string
	CountryName string
}

func (city City) returnCityName() string {
	return city.CityName
}

func (country City) return_Country_With_City_Name() string {
	return country.CountryName
}

func calculateCityAndCountyExist(city City) bool {
	Return_City := city.returnCityName()
	Return_Country := city.CountryName
	
	if Return_City == Return_Country {
		return true
	} else {
		return false
	}

}




