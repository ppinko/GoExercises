package main

import "fmt"

const OvenTime int = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	return 2 * numberOfLayers
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

func main() {
	fmt.Println(RemainingOvenTime(10))
	fmt.Println(PreparationTime(2))
	fmt.Println(ElapsedTime(3, 20))
}
