// Package weather forecast the current weather condition of various cities in Goblinocus.
package weather

var (
    // CurrentCondition represents the weather condition in a particular city.
	CurrentCondition string
    // CurrentLocation represents the city whose weather is currently forecasted.
	CurrentLocation  string
)

// Forecast returns a string value that is equal to the current weather condition of a particular city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
