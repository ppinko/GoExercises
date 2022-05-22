package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time
func SeedWithTime() {
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20
func RollADie() int {
	return 1 + rand.Intn(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0
func GenerateWandEnergy() float64 {
	return 12.0 * rand.Float64()
}

// ShuffleAnimals returns a slice with all eight animal strings in random order
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}

func main() {
	fmt.Println()
	SeedWithTime()
	fmt.Println(RollADie())           // a random int, 1 <= d <= 20
	fmt.Println(GenerateWandEnergy()) // a random float64, 0.0 <= f < 12.0
	fmt.Println(ShuffleAnimals())
}
