// Package weather refers to the temperature of the weather.
package weather

// CurrentCondition Defines the current situation.
var CurrentCondition string

// CurrentLocation Specifies the current location.
var CurrentLocation string

// Forecast but It predicts the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
