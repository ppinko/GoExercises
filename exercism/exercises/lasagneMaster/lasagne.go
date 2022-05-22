package main

import "fmt"

// https://exercism.org/tracks/go/exercises/lasagna-master

func PreparationTime(layers []string, avgTimePerLayer int) int {
	if avgTimePerLayer == 0 {
		avgTimePerLayer = 2
	}
	return len(layers) * avgTimePerLayer
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, val := range layers {
		if val == "noodles" {
			noodles += 50
		} else if val == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	amount := make([]float64, len(quantities))
	copy(amount, quantities)
	scale := float64(portions) * 0.5
	for i, _ := range amount {
		amount[i] = amount[i] * scale
	}
	return amount
}

func main() {
	fmt.Println()

	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3)) // => 18
	fmt.Println(PreparationTime(layers, 0)) // => 12

	fmt.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})) // => 100, 0.4

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	AddSecretIngredient(friendsList, myList)

	fmt.Println(myList)
	// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}

	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Println(scaledQuantities)
	// => []float64{ 2.4, 7.2, 21 }
}
