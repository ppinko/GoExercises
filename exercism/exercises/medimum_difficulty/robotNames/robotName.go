package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

// https://exercism.org/tracks/go/exercises/robot-name

var RobotNames map[string]bool = make(map[string]bool)

const Alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const AlphabetLength int = 26
const NumDigits int = 10
const NumPermutations int = 26 * 26 * 10 * 10 * 10

type Robot struct {
	RobotName string
}

func (r *Robot) Name() (string, error) {
	if r.RobotName != "" {
		return r.RobotName, nil
	}

	for i := 0; i < NumPermutations; i++ {
		char1 := rand.Intn(AlphabetLength)
		char2 := rand.Intn(AlphabetLength)
		num1 := rand.Intn(NumDigits)
		num2 := rand.Intn(NumDigits)
		num3 := rand.Intn(NumDigits)

		name := string(Alphabet[char1]) + string(Alphabet[char2]) + strconv.Itoa(num1) + strconv.Itoa(num2) + strconv.Itoa(num3)
		_, exists := RobotNames[name]

		if !exists {
			RobotNames[name] = true
			r.RobotName = name
			return name, nil
		}
	}
	return "", errors.New("no more names available")
}

func (r *Robot) Reset() {
	oldName := r.RobotName
	delete(RobotNames, oldName)
	r.RobotName = ""
}

func main() {
	fmt.Println()
}
