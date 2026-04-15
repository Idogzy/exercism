package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * successRate/100
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    result := CalculateWorkingCarsPerHour(productionRate, successRate) / 60
    return int(result)
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    groupOfTen := carsCount / 10
    remainder := carsCount % 10
    result := groupOfTen * 95000 + remainder * 10000
    return uint(result)
	panic("CalculateCost not implemented")
}
