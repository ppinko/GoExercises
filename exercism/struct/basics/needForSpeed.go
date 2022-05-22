package main

import "fmt"

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func NewCar(speed, batterDain int) Car {
	return Car{100, batterDain, speed, 0}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{distance: distance}
}

func Drive(car Car) Car {
	if car.battery-car.batteryDrain > 0 {
		return Car{
			battery:      car.battery - car.batteryDrain,
			batteryDrain: car.batteryDrain,
			speed:        car.speed,
			distance:     car.distance + car.speed,
		}
	} else {
		return car
	}
}

func CanFinish(car Car, track Track) bool {
	// initDistance := car.distance
	distance := car.battery * car.speed / car.batteryDrain
	if distance >= track.distance {
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
	// Output: Car{speed: 5, batteryDrain: 2, battery:100, distance: 0}

	// distance := 800
	// raceTrack := NewTrack(distance)
	// Output: Track{distance: 800}

	car = Drive(car)
	// Output: Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}

	car2 := NewCar(speed, batteryDrain)
	distance2 := 100
	raceTrack2 := NewTrack(distance2)
	fmt.Println(CanFinish(car2, raceTrack2))
	// Output: true
}
