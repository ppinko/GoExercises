package main

import (
	"fmt"
	"sort"
)

// https://exercism.org/tracks/go/exercises/sum-of-multiples

func SumMultiples(limit int, divisors ...int) int {
	sort.Slice(divisors, func(i, j int) bool { return i > j })
	sum := 0

	// use only positive integers
	for i, v := range divisors {
		if v >= 1 {
			divisors = divisors[i:]
			break
		}
	}

	// handle special cases
	if len(divisors) == 0 {
		return sum
	} else if divisors[0] == 1 {
		return ArithmeticSum(1, limit)
	}

	// handle typical cases
	for i := 2; i < limit; i++ {
		if IsDividible(i, divisors) {
			sum += i
		}
	}
	return sum
}

func IsDividible(divisor int, nums []int) bool {
	for _, v := range nums {
		if divisor%v == 0 {
			return true
		}
	}

	return false
}

func ArithmeticSum(divisor, limit int) int {
	sum := 0
	if limit <= divisor {
		return sum
	}
	max := limit / divisor
	if limit%divisor == 0 {
		max--
	}
	sum = divisor * (1 + max) * max / 2
	return sum
}

func GreatestCommonDivisor(a, b int) int {
	for a%b != 0 {
		temp := a % b
		a = b
		b = temp
	}
	return b
}

func main() {
	fmt.Println(SumMultiples(4, 3, 0))
	fmt.Printf("Sum of multiplies of 3 and 5 lower than 20 = %d",
		SumMultiples(100, 3, 5))
}
