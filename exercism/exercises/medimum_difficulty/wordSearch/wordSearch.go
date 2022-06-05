package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func FindWord(puzzle []string, word string, i, j int) (bool, [2][2]int) {
	rows := len(puzzle)

	var position [2][2]int
	increment := utf8.RuneCountInString(word) - 1

	hRight := puzzle[i][j:]
	idx := strings.Index(hRight, word)
	if idx != -1 {
		position[0][1] = i
		position[0][0] = j + idx
		position[1][1] = i
		position[1][0] = j + idx + increment
		return true, position
	}

	hLeft := string(puzzle[i][j])
	for k := j - 1; k >= 0; k-- {
		hLeft += string(puzzle[i][k])
	}
	idx = strings.Index(hLeft, word)
	if idx != -1 {
		position[0][1] = i
		position[0][0] = j - idx
		position[1][1] = i
		position[1][0] = j - idx - increment
		return true, position
	}

	vDown := string(puzzle[i][j])
	for k := i + 1; k < rows; k++ {
		if j >= utf8.RuneCountInString(puzzle[k]) {
			break
		}
		vDown += string(puzzle[k][j])
	}
	idx = strings.Index(vDown, word)
	if idx != -1 {
		position[0][1] = i + idx
		position[0][0] = j
		position[1][1] = i + idx + increment
		position[1][0] = j
		return true, position
	}

	vUp := string(puzzle[i][j])
	for k := i - 1; k >= 0; k-- {
		if j >= utf8.RuneCountInString(puzzle[k]) {
			break
		}
		vUp += string(puzzle[k][j])
	}
	idx = strings.Index(vUp, word)
	if idx != -1 {
		position[0][1] = i - idx
		position[0][0] = j
		position[1][1] = i - idx - increment
		position[1][0] = j
		return true, position
	}

	diagUpLeft := string(puzzle[i][j])
	m := j - 1
	for k := i - 1; k >= 0 && m >= 0; k-- {
		if m >= utf8.RuneCountInString(puzzle[k]) || m < 0 {
			break
		}
		diagUpLeft += string(puzzle[k][m])
		m--
	}
	idx = strings.Index(diagUpLeft, word)
	if idx != -1 {
		position[0][1] = i - idx
		position[0][0] = j - idx
		position[1][1] = i - idx - increment
		position[1][0] = j - idx - increment
		return true, position
	}

	diagUpRight := string(puzzle[i][j])
	m = j + 1
	for k := i - 1; k >= 0; k-- {
		if m >= utf8.RuneCountInString(puzzle[k]) {
			break
		}
		diagUpRight += string(puzzle[k][m])
		m++
	}
	idx = strings.Index(diagUpRight, word)
	if idx != -1 {
		position[0][1] = i - idx
		position[0][0] = j + idx
		position[1][1] = i - idx - increment
		position[1][0] = j + idx + increment
		return true, position
	}

	diagDownLeft := string(puzzle[i][j])
	m = j - 1
	for k := i + 1; k < rows; k++ {
		if m >= utf8.RuneCountInString(puzzle[k]) || m < 0 {
			break
		}
		diagDownLeft += string(puzzle[k][m])
		m--
	}
	idx = strings.Index(diagDownLeft, word)
	if idx != -1 {
		position[0][1] = i + idx
		position[0][0] = j - idx
		position[1][1] = i + idx + increment
		position[1][0] = j - idx - increment
		return true, position
	}

	diagDownRight := string(puzzle[i][j])
	m = j + 1
	for k := i + 1; k < rows; k++ {
		if m >= utf8.RuneCountInString(puzzle[k]) {
			break
		}
		diagDownRight += string(puzzle[k][m])
		m++
	}
	idx = strings.Index(diagDownRight, word)
	if idx != -1 {
		position[0][1] = i + idx
		position[0][0] = j + idx
		position[1][1] = i + idx + increment
		position[1][0] = j + idx + increment
		return true, position
	}

	return false, position
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	solved := make(map[string][2][2]int)
	if len(words) == 0 {
		return solved, errors.New("no words to look for")
	} else if len(puzzle) == 0 {
		return solved, errors.New("empty puzzle")
	}

	for _, word := range words {
		char := string(word[0])
		exit := false
		for i, r := range puzzle {
			if exit {
				break
			}
			for j, c := range r {
				if string(c) == char {
					found, position := FindWord(puzzle, word, i, j)
					if found {
						solved[word] = position
						exit = true
						break
					}
				}
			}

			if i == len(puzzle)-1 && !exit {
				return solved, errors.New("word not found")
			}
		}
	}
	return solved, nil
}

func main() {
	puzzle := []string{
		"jefblpepre",
		"camdcimgtc",
		"oivokprjsm",
		"pbwasqroua",
		"rixilelhrs",
		"wolcqlirpc",
		"screeaumgr",
		"alxhpburyi",
		"jalaycalmp",
		"clojurermt",
	}

	words := []string{"ecmascript", "rust"}

	solved, _ := Solve(words, puzzle)
	fmt.Println(solved)
}
