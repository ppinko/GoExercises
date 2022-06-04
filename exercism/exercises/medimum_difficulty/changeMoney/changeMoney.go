package main

import (
	"errors"
	"fmt"
	"sort"
)

// https://exercism.org/tracks/go/exercises/change

func LastSolution(solution []int) bool {
	if len(solution) <= 1 {
		return true
	}

	for i := 0; i < len(solution)-1; i++ {
		if solution[i] != 0 {
			return false
		}
	}
	return true
}

func FillRest(idx int, solution []int) {
	for idx < len(solution) {
		solution[idx] = 0
		idx++
	}
}

func FirstTry(target int, coins, solution []int) bool {
	for i := 0; i < len(coins); i++ {
		if target/coins[i] == 0 {
			solution[i] = 0
		} else {
			amount := target / coins[i]
			solution[i] = amount
			target -= coins[i] * amount
			if target == 0 {
				FillRest(i+1, solution)
				return true
			}
		}
	}
	return false
}

func SumToIdx(target, idx int, coins, solution []int) int {
	for i := 0; i < idx; i++ {
		target -= coins[i] * solution[i]
	}
	return target
}

func SumSlice(solution []int) int {
	sum := 0
	for _, v := range solution {
		sum += v
	}
	return sum
}

func Next(target int, coins, solution []int) bool {
	// find starting index
	idxLeft := 0
	for i, v := range solution {
		if v != 0 {
			idxLeft = i
			break
		}
	}

	idxRight := idxLeft
	for j := len(solution) - 2; j >= idxLeft; j-- {
		if solution[j] != 0 {
			idxRight = j
			break
		}
	}

	if idxRight != idxLeft {
		target = SumToIdx(target, idxRight, coins, solution)
	}
	solution[idxRight] = solution[idxRight] - 1
	target -= coins[idxRight] * solution[idxRight]
	for i := idxRight + 1; i < len(solution); i++ {
		amount := target / coins[i]
		if amount == 0 {
			solution[i] = 0
		} else {
			solution[i] = amount
			target -= coins[i] * amount
			if target == 0 {
				FillRest(i+1, solution)
				return true
			}
		}
	}

	return false
}

func FinalSolution(coins, solution []int) []int {
	var result []int
	for i := 0; i < len(solution); i++ {
		for solution[i] > 0 {
			result = append(result, coins[i])
			solution[i]--
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

func Change(coins []int, target int) ([]int, error) {
	result := make([]int, len(coins))
	best := make([]int, len(coins))
	bestSum := 0
	initialized := false
	sort.Slice(coins, func(i, j int) bool { return coins[i] > coins[j] })

	// check corner cases
	if len(coins) < 0 {
		return result, errors.New("no coins")
	} else if target == 0 {
		return make([]int, 0), nil
	} else if coins[len(coins)-1] > target || target < 0 {
		return make([]int, 0), errors.New("change cannot be smaller than the smallest coin")
	}

	successful := FirstTry(target, coins, result)
	if successful {
		initialized = true
		copy(best, result)
		bestSum = SumSlice(best)
	}

	for {
		isValid := Next(target, coins, result)
		if isValid {
			if !initialized {
				initialized = true
				copy(best, result)
				bestSum = SumSlice(best)
			} else {
				tempSum := SumSlice(result)
				if tempSum < bestSum {
					copy(best, result)
					bestSum = tempSum
				}
			}
		}
		if LastSolution(result) {
			break
		}
	}

	if !initialized {
		return nil, errors.New("no valid solution")
	} else {
		sol := FinalSolution(coins, best)
		return sol, nil
	}
}

func main() {
	change, _ := Change([]int{50, 20, 10, 5, 2}, 21)
	fmt.Println(change)
}
