package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// https://exercism.org/tracks/go/exercises/queen-attack

type position [2]int

func ValidateInput(field string) (position, error) {
	var pos position
	if utf8.RuneCountInString(field) != 2 {
		return pos, errors.New("wrong input")
	}
	for i, v := range field {
		num := 0
		if i == 0 {
			num = int(v - 'a')
		} else {
			num = int(v - '1')
		}
		if num < 0 || num > 7 {
			return pos, errors.New("wrong input")
		} else {
			pos[i] = num
		}
	}
	return pos, nil
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	// parse and validate input
	pos1, err1 := ValidateInput(whitePosition)
	pos2, err2 := ValidateInput(blackPosition)
	if err1 != nil || err2 != nil || whitePosition == blackPosition {
		return false, errors.New("wrong input")
	}

	// same row or column
	if pos1[0] == pos2[0] || pos1[1] == pos2[1] {
		return true, nil
	}

	// curve slope must be 45deg or -45deg (a=1 or a=-1)
	remainder := (pos2[1] - pos1[1]) % (pos2[0] - pos1[0])
	a := (pos2[1] - pos1[1]) / (pos2[0] - pos1[0])
	if remainder == 0 && (a == 1 || a == -1) {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	v, _ := CanQueenAttack("c5", "f2")
	fmt.Printf("Can they attack each other ? %t", v)
}
