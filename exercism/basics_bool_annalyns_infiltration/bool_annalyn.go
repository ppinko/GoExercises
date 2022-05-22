package main

import "fmt"

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (petDogIsPresent && !archerIsAwake) || (!knightIsAwake && !archerIsAwake && prisonerIsAwake)
}

func main() {
	// can be fast attacked
	var knightIsAwake bool = true
	fmt.Println(CanFastAttack(knightIsAwake))

	// can be spied
	knightIsAwake = false
	var archerIsAwake, prisonerIsAwake bool = true, false
	fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))

	// can signal prisoner
	archerIsAwake, prisonerIsAwake = false, true
	fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))

	// can free prisoner
	knightIsAwake = false
	archerIsAwake = true
	prisonerIsAwake = false
	petDogIsPresent := false
	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
}
