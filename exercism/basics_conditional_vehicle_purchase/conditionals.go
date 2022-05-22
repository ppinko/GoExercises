package main

import "fmt"

func NeedsLicense(kind string) bool {
	if kind == "car" {
		return true
	} else if kind == "truck" {
		return true
	} else {
		return false
	}
}

func ChooseVehicle(option1, option2 string) string {
	var vehicle string
	if option1 > option2 {
		vehicle = option2
	} else {
		vehicle = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", vehicle)
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	var currentPrice float64
	if age < 3.0 {
		currentPrice = 0.8 * originalPrice
	} else if age < 10.0 {
		currentPrice = 0.7 * originalPrice
	} else {
		currentPrice = 0.5 * originalPrice
	}
	return currentPrice
}

func main() {
	needLicense := NeedsLicense("car")
	fmt.Println(needLicense)

	vehicle := ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
	fmt.Println(vehicle)

	fmt.Println(CalculateResellPrice(1000, 5))
}
