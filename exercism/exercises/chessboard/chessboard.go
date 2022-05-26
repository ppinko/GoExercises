package main

import "fmt"

// https://exercism.org/tracks/go/exercises/chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank [8]bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	counter := 0
	val, exists := cb[rank]
	if exists {
		for _, v := range val {
			if v {
				counter++
			}
		}
	}
	return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	counter := 0
	if file < 1 || file > 8 {
		return counter
	}
	file--
	for _, rank := range cb {
		if rank[file] {
			counter++
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	counter := 0
	for _, val := range cb {
		for range val {
			counter++
		}
	}
	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, val := range cb {
		for _, occupied := range val {
			if occupied {
				counter++
			}
		}
	}
	return counter
}

func main() {
	fmt.Println()
}
