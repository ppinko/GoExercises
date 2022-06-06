package main

// https://exercism.org/tracks/go/exercises/matching-brackets

func Bracket(input string) bool {
	// maps
	closingBrackets := map[rune]rune{'}': '{', ']': '[', ')': '('}
	openingBrackets := map[rune]bool{'{': true, '[': true, '(': true}
	// LIFO
	var lifo []rune
	for _, r := range input {
		if _, exists := openingBrackets[r]; exists {
			lifo = append(lifo, r)
		} else if v, exists := closingBrackets[r]; exists {
			if len(lifo) > 0 && lifo[len(lifo)-1] == v {
				lifo = lifo[:len(lifo)-1]
			} else {
				return false
			}
		}
	}
	return len(lifo) == 0
}

func main() {
}
