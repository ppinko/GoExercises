package main

import "fmt"

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

func PrependItems(slice []int, value ...int) []int {
	if len(value) == 0 {
		return slice
	}

	var prepend []int
	for _, val := range value {
		prepend = append(prepend, val)
	}
	slice = append(prepend, slice...)
	return slice
}

func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	start := slice[:index]
	end := slice[index+1:]
	slice = append(start, end...)
	return slice
}

func main() {
	fmt.Println()
	fmt.Printf("%v", FavoriteCards())

	fmt.Println()
	cards := FavoriteCards()
	fmt.Println(cards)

	fmt.Println()
	card := GetItem([]int{1, 2, 4, 1}, 2)
	fmt.Println(card)

	fmt.Println()
	index := 2
	newCard := 6
	cards2 := SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards2)

	slice := []int{3, 2, 6, 4, 8}
	cards3 := PrependItems(slice, 5, 1)
	fmt.Println(cards3)

	cards4 := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println(cards4)
}
