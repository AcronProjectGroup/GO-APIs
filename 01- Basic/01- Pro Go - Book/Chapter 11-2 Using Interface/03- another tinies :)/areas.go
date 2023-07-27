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

func (countryArea Area) getCountryArea() float64 {
	return countryArea.CountryArea
}


