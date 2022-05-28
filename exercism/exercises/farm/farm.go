package main

import (
	"errors"
	"fmt"
)

// https://exercism.org/tracks/go/exercises/the-farm

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	Cows int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	} else if cows < 0 {
		s := SillyNephewError{Cows: cows}
		return 0.0, &s
	}

	val, err := weightFodder.FodderAmount()
	if err == nil {
		return val / float64(cows), err
	} else if err == ErrScaleMalfunction && val > 0.0 {
		return 2 * val / float64(cows), nil
	} else if (err == ErrScaleMalfunction || err == nil) && val < 0.0 {
		return 0.0, errors.New("negative fodder")
	} else {
		return 0.0, err
	}
}

func main() {
	fmt.Println()

	// twentyFodderNoError says there are 20.0 fodder
	fodder, err := DivideFood(twentyFodderNoError, 10)
	// fodder == 2.0
	// err == nil
}
