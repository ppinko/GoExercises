package main

import "fmt"

func CalculateWorkingCarsPerHour(numberOfCarsPerHour, successRate int) float64 {
	return float64(numberOfCarsPerHour) * float64(successRate) / 100.0
}

func CalculateWorkingCarsPerMinute(numberOfCarsPerHour, successRate int) int {
	return int(CalculateWorkingCarsPerHour(numberOfCarsPerHour, successRate) / 60)
}

func CalculateCost(numberOfCars uint) uint {
	var groups uint = numberOfCars / 10
	var remaining uint = numberOfCars % 10
	return groups*95000 + remaining*10000
}

func main() {
	rate := CalculateWorkingCarsPerHour(1547, 90)
	fmt.Println(rate)

	rateMin := CalculateWorkingCarsPerMinute(1105, 90)
	fmt.Println(rateMin)

	cost := CalculateCost(37)
	fmt.Println(cost)
}
