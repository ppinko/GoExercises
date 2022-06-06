package main

// https://exercism.org/tracks/go/exercises/two-bucket

import (
	"errors"
)

type Buckets struct {
	CapA    int
	CapB    int
	AmountA int
	AmountB int
	Actions int
}

func (b *Buckets) Validate(goalAmount int) bool {
	return b.AmountA == goalAmount || b.AmountB == goalAmount
}

func (b *Buckets) Output(goalAmount int) (string, int, int, error) {
	target := ""
	amount := 0
	if b.AmountA == goalAmount {
		target = "one"
		amount = b.AmountB
	} else {
		target = "two"
		amount = b.AmountA
	}
	return target, b.Actions, amount, nil
}

func (b *Buckets) InMap(sols map[Buckets]bool) bool {
	if b.AmountA == 0 && b.AmountB == 0 {
		return false
	}

	for key, _ := range sols {
		if key.AmountA == b.AmountA && key.AmountB == b.AmountB {
			return false
		}
	}
	sols[*b] = true
	return true
}

func (b *Buckets) Next(goalAmount int, startingBucket string, sols map[Buckets]bool) []Buckets {
	var buckets []Buckets
	if b.CapA == b.AmountA && b.CapB == b.AmountB {
		// 1st case, both buckets are full
		if startingBucket == "one" {
			temp := Buckets{b.CapA, b.CapB, b.AmountA, 0, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else {
			temp := Buckets{b.CapA, b.CapB, 0, b.AmountB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
	} else if b.CapA == b.AmountA && b.AmountB == 0 {
		// 2nd case, A full, B empty
		if b.CapB >= b.CapA {
			temp := Buckets{b.CapA, b.CapB, 0, b.AmountA, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if b.CapB < b.CapA {
			temp := Buckets{b.CapA, b.CapB, b.AmountA - b.CapB, b.CapB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
		temp2 := Buckets{b.CapA, b.CapB, b.AmountA, b.CapB, b.Actions + 1}
		if temp2.InMap(sols) {
			buckets = append(buckets, temp2)
		}
	} else if b.AmountA > 0 && b.AmountB == 0 {
		// 2nd case, A full, B empty
		if b.CapB > b.AmountA {
			temp := Buckets{b.CapA, b.CapB, 0, b.AmountA, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if b.CapB < b.AmountA {
			temp := Buckets{b.CapA, b.CapB, b.AmountA - b.CapB, b.CapB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if b.AmountA == b.CapB && startingBucket == "two" {
			temp := Buckets{b.CapA, b.CapB, 0, b.CapB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
		temp2 := Buckets{b.CapA, b.CapB, b.AmountA, b.CapB, b.Actions + 1}
		if temp2.InMap(sols) {
			buckets = append(buckets, temp2)
		}
	} else if b.CapB == b.AmountB && b.AmountA == 0 {
		// 3rd case, B full, A empty
		if b.CapA >= b.CapB {
			temp := Buckets{b.CapA, b.CapB, b.AmountB, 0, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if b.CapA < b.CapB {
			temp := Buckets{b.CapA, b.CapB, b.CapA, b.AmountB - b.CapA, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
		temp2 := Buckets{b.CapA, b.CapB, b.CapA, b.CapB, b.Actions + 1}
		if temp2.InMap(sols) {
			buckets = append(buckets, temp2)
		}
	} else if b.AmountB > 0 && b.AmountA == 0 {
		if b.CapA > b.AmountB {
			temp := Buckets{b.CapA, b.CapB, b.AmountB, 0, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if b.CapA < b.AmountB {
			temp := Buckets{b.CapA, b.CapB, b.CapA, b.AmountB - b.CapA, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if b.AmountB == b.CapA && startingBucket == "one" {
			temp := Buckets{b.CapA, b.CapB, b.CapA, 0, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
		temp2 := Buckets{b.CapA, b.CapB, b.CapA, b.AmountB, b.Actions + 1}
		if temp2.InMap(sols) {
			buckets = append(buckets, temp2)
		}
	} else if b.CapA == b.AmountA && b.AmountB != 0 {
		// 4th case, A full, B non-empty
		temp2 := Buckets{b.CapA, b.CapB, 0, b.AmountB, b.Actions + 1}
		if temp2.InMap(sols) {
			buckets = append(buckets, temp2)
		}
		if startingBucket == "one" {
			temp := Buckets{b.CapA, b.CapB, b.AmountA, 0, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
		temp3 := Buckets{b.CapA, b.CapB, b.AmountA, b.CapB, b.Actions + 1}
		if temp3.InMap(sols) {
			buckets = append(buckets, temp3)
		}
		if b.AmountA > b.CapB-b.AmountB {
			temp := Buckets{b.CapA, b.CapB, b.AmountA - (b.CapB - b.AmountB), b.CapB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if startingBucket == "two" || b.AmountA+b.AmountB != b.CapB {
			temp := Buckets{b.CapA, b.CapB, 0, b.AmountA + b.AmountB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
	} else if b.CapB == b.AmountB && b.AmountA != 0 {
		// 4th case, B full, A non-empty
		if startingBucket == "two" {
			temp := Buckets{b.CapA, b.CapB, 0, b.AmountB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
		temp2 := Buckets{b.CapA, b.CapB, b.AmountA, 0, b.Actions + 1}
		if temp2.InMap(sols) {
			buckets = append(buckets, temp2)
		}
		temp3 := Buckets{b.CapA, b.CapB, b.CapA, b.AmountB, b.Actions + 1}
		if temp3.InMap(sols) {
			buckets = append(buckets, temp3)
		}
		if b.AmountB > b.CapA-b.AmountA {
			temp := Buckets{b.CapA, b.CapB, b.CapA, b.AmountB - (b.CapA - b.AmountA), b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if startingBucket == "one" || b.AmountA+b.AmountB != b.CapA {
			temp := Buckets{b.CapA, b.CapB, b.AmountA + b.AmountB, 0, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
	} else {
		temp2 := Buckets{b.CapA, b.CapB, 0, b.AmountB, b.Actions + 1}
		if temp2.InMap(sols) {
			buckets = append(buckets, temp2)
		}
		temp3 := Buckets{b.CapA, b.CapB, b.AmountA, 0, b.Actions + 1}
		if temp3.InMap(sols) {
			buckets = append(buckets, temp3)
		}
		temp4 := Buckets{b.CapA, b.CapB, b.CapA, b.AmountB, b.Actions + 1}
		if temp4.InMap(sols) {
			buckets = append(buckets, temp4)
		}
		temp5 := Buckets{b.CapA, b.CapB, b.AmountA, b.CapB, b.Actions + 1}
		if temp5.InMap(sols) {
			buckets = append(buckets, temp5)
		}
		if b.AmountA > b.CapB-b.AmountB {
			temp := Buckets{b.CapA, b.CapB, b.AmountA - (b.CapB - b.AmountB), b.CapB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if startingBucket == "two" || b.AmountA+b.AmountB != b.CapB {
			temp := Buckets{b.CapA, b.CapB, 0, b.AmountA + b.AmountB, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
		if b.AmountB > b.CapA-b.AmountA {
			temp := Buckets{b.CapA, b.CapB, b.CapA, b.AmountB - (b.CapA - b.AmountA), b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		} else if startingBucket == "one" || b.AmountA+b.AmountB != b.CapA {
			temp := Buckets{b.CapA, b.CapB, b.AmountA + b.AmountB, 0, b.Actions + 1}
			if temp.InMap(sols) {
				buckets = append(buckets, temp)
			}
		}
	}
	return buckets
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 || goalAmount <= 0 || (startBucket != "one" && startBucket != "two") {
		return "", 0, 0, errors.New("invalid input")
	}

	var buckets []Buckets
	initBucket := Buckets{sizeBucketOne, sizeBucketTwo, 0, 0, 1}
	if startBucket == "one" {
		initBucket.AmountA = initBucket.CapA
	} else {
		initBucket.AmountB = initBucket.CapB
	}

	solutions := make(map[Buckets]bool)
	solutions[initBucket] = true
	buckets = append(buckets, initBucket)

	// main loop
	i := 0
	for {
		if buckets[i].Validate(goalAmount) {
			return buckets[i].Output(goalAmount)
		} else {
			toAdd := buckets[i].Next(goalAmount, startBucket, solutions)
			buckets = append(buckets, toAdd...)
		}
		i++
		if i >= len(buckets) {
			break
		}
	}

	return "", 0, 0, nil
}

func main() {
	// _, _, _, _ = Solve(3, 5, 1, "one")
	// _, _, _, _ = Solve(3, 5, 1, "two")
	_, _, _, _ = Solve(7, 11, 2, "one")
}
