package main

import "fmt"

func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for _, v := range birdsPerDay {
		sum += v
	}
	return sum
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	subslice := birdsPerDay[7*(week-1) : 7*week]
	return TotalBirdCount(subslice)
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}

func main() {
	fmt.Println()

	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(TotalBirdCount(birdsPerDay)) // => 34

	fmt.Println(BirdsInWeek(birdsPerDay, 2)) // => 12

	birdsPerDay2 := []int{2, 5, 0, 7, 4, 1}
	fmt.Println(FixBirdCountLog(birdsPerDay2)) // => [3 5 1 7 5 1]
}
