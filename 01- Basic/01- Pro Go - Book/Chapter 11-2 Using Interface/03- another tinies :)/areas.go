package main

type Area struct {
	CountryName string
	CityName    string
	CountryArea float64
}

func (countryName Area) returnCountryName() string {
	return countryName.CountryName
}

func (cityName Area) returnCityName() string {
	return cityName.CityName
}

func calculateCityAndCountyExist(area Area) float64 {
	Return_City := area.returnCityName()
	Return_Country := area.CountryName

	if Return_City == Return_Country {
		return area.CountryArea
	} else {
		return 0
	}

}
