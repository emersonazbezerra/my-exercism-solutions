// Package weather provides the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the actual weather condition of the location.
var CurrentCondition string
// CurrentLocation represents the name of the current location.
var CurrentLocation string

// Forecast returns the weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
