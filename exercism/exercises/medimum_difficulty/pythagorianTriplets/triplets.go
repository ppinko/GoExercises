package main

import (
	"fmt"
	"sort"
)

// task - https://exercism.org/tracks/go/exercises/pythagorean-triplet

// theory - https://www.cuemath.com/geometry/pythagorean-triples/

type Triplet [3]int

/*
a = m^2 - n^2
b = 2mn
c = m^2 + n^2


*/

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var result []Triplet
	// even values
	evenPossible, oddPossible := true, true
	i := 2
	for evenPossible || oddPossible {
		middle := 0
		upper := 0
		bottom := i
		if evenPossible && i%2 == 0 {
			core := i * i / 4
			upper = core + 1
			middle = core - 1
			if upper > max {
				evenPossible = false
			}
		} else if oddPossible { // odd values
			upper = (i*i + 1) / 2
			middle = upper - 1
			if upper > max {
				oddPossible = false
			}
		}

		if i < min && upper < max && upper != 0 {
			scale := max / upper
			if scale != 0 {
				upper *= scale
				bottom *= scale
				middle *= scale
			}
		}

		if upper <= max && bottom >= min && bottom < middle && middle < upper {
			triplet := Triplet{bottom, middle, upper}
			toAppend := true
			for _, v := range result {
				if v == triplet {
					toAppend = false
				}
			}
			if toAppend {
				result = append(result, Triplet{bottom, middle, upper})
			}
		}

		i++
	}
	return result
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var result []Triplet
	// even values
	max := p / 2
	min := 3
	evenPossible, oddPossible := true, true
	i := 2
	for evenPossible || oddPossible {
		middle := 0
		upper := 0
		bottom := i
		if evenPossible && i%2 == 0 {
			core := i * i / 4
			upper = core + 1
			middle = core - 1
			if upper > max {
				evenPossible = false
			}
		} else if oddPossible { // odd values
			upper = (i*i + 1) / 2
			middle = upper - 1
			if upper > max {
				oddPossible = false
			}
		}

		if i < min && upper < max && upper != 0 {
			scale := max / upper
			if scale != 0 {
				upper *= scale
				bottom *= scale
				middle *= scale
			}
		}

		if upper <= max && bottom >= min && bottom < middle && middle < upper {
			toAppend := false
			triplet := Triplet{bottom, middle, upper}
			subSum := triplet[0] + triplet[1] + triplet[2]
			if subSum == p {
				toAppend = true
			} else if subSum < p && p%subSum == 0 {
				scale := p / subSum
				triplet = Triplet{scale * bottom, scale * middle, scale * upper}
				toAppend = true
			}

			if toAppend {
				for _, v := range result {
					if v == triplet {
						toAppend = false
					}
				}
			}

			if toAppend {
				result = append(result, triplet)
			}
		}

		i++
	}
	sort.Slice(result, func(i, j int) bool { return result[i][0] < result[j][0] })
	return result
}

func main() {
	fmt.Println(Range(1, 10))
	fmt.Println(Range(11, 20))
	fmt.Println(Sum(180))
	fmt.Println(Sum(1000))
}
