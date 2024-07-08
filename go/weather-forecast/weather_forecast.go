// Package weather provides a function that returns a weather forecast for a given city.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current location being forecasted.
var CurrentLocation string

// Forecast returns a weather forecast for a given city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
