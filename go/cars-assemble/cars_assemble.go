package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    minutesInHour := 60.0
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / minutesInHour)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const (
        tenCarsCost = 95000
    	individualCarCost = 10000
	)

    groupsOfTenCarsToBeProduced := carsCount / 10
    individualCarsToBeProduced := carsCount % 10

    return uint((tenCarsCost * groupsOfTenCarsToBeProduced) + (individualCarsToBeProduced * individualCarCost))
}
