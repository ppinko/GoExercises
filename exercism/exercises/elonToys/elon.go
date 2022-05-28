package main

import "fmt"

// https://exercism.org/tracks/go/exercises/elons-toys/edit

// car struct
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
}

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	if car.battery/car.batteryDrain >= trackDistance/car.speed {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println()

	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	car.Drive()
	// car is now Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}

	fmt.Println(car.DisplayDistance())

	fmt.Println(car.DisplayBattery())

	trackDistance := 100

	car.CanFinish(trackDistance)
}
