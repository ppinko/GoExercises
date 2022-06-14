package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Transpose(input []string) []string {
	if len(input) == 0 {
		return input
	}

	// create copy of input
	cp := make([]string, len(input))
	copy(cp, input)

	// find maximum row length
	maxLength := 0
	for _, str := range cp {
		if utf8.RuneCountInString(str) > maxLength {
			maxLength = utf8.RuneCountInString(str)
		}
	}

	// make each string of equal length by appending whitespaces
	for i := 0; i < len(cp); i++ {
		toAppend := maxLength - utf8.RuneCountInString(cp[i])
		if maxLength != 0 {
			cp[i] += strings.Repeat("%", toAppend)
		}
	}

	// due to internal representation of strings use [][]rune instead
	base := make([][]rune, len(cp))
	for i, str := range cp {
		base[i] = []rune(str)
	}

	// transpose slice of slice of runes
	transpose := make([][]rune, maxLength)
	for i := 0; i < maxLength; i++ {
		row := make([]rune, len(cp))
		for j := 0; j < len(cp); j++ {
			row[j] = base[j][i]
		}
		transpose[i] = row
	}

	// create slice of strings and remove trailing whitespaces
	answer := make([]string, maxLength)
	for i := 0; i < maxLength; i++ {
		str := string(transpose[i])
		str = strings.TrimRight(str, "%")
		str = strings.Replace(str, "%", " ", -1)
		answer[i] = str
	}

	return answer
}

func main() {
	s := []string{"ABC", "123"}
	ret := Transpose(s)
	fmt.Println(ret)
}
