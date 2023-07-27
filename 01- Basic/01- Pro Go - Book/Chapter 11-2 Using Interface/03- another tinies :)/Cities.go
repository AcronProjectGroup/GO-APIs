
package main

type City struct {
	CountryName string
	CityName string
}

func (countryName City) returnCountryName() string {
	return countryName.CountryName
}

func (cityName City) returnCityName() string {
	return cityName.CityName
}

func (city City)getCountryArea() float64 {
	return 0
}





