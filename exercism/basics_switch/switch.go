package main

// https://exercism.org/tracks/go/exercises/blackjack

import "fmt"

func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	myCards := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)

	switch {
	case myCards == 22:
		return "P"
	case myCards == 21 && dealer >= 10:
		return "S"
	case myCards == 21 && dealer < 10:
		return "W"
	case myCards >= 17:
		return "S"
	case myCards >= 12 && dealer < 7:
		return "S"
	case myCards >= 12 && dealer >= 7:
		return "H"
	default:
		return "H"
	}
}

func main() {
	value := ParseCard("ace")
	fmt.Println(value) // Output: 11

	fmt.Println("Success")
}
