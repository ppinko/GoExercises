package main

import (
	"fmt"
	"strconv"
)

// https://exercism.org/tracks/go/exercises/run-length-encoding

func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}
	encoded := ""
	first, last := 0, 0
	var char rune
	for i, v := range input {
		if i == 0 {
			first = i
			last = i
			char = v
			continue
		}
		if char == v {
			continue
		} else {
			last = i
			repeats := last - first
			if repeats > 1 {
				encoded += strconv.Itoa(repeats)
			}
			encoded += string(char)

			// set initial values for next
			first = i
			char = v
		}
	}
	// handle last value
	last = len(input)
	repeats := last - first
	if repeats > 1 {
		encoded += strconv.Itoa(repeats)
	}
	encoded += string(char)

	return encoded
}

func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return ""
	}
	decoded := ""
	num := ""
	for _, v := range input {
		if int('9'-v) >= 0 && int('9'-v) <= 9 {
			num += string(v)
		} else {
			if len(num) == 0 {
				decoded += string(v)
			} else {
				converter, _ := strconv.Atoi(num)
				for i := 0; i < converter; i++ {
					decoded += string(v)
				}
				num = ""
			}
		}
	}

	return decoded
}

func main() {
	fmt.Println(RunLengthEncode("WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"))
	// output: "12WB12W3B24WB"
	fmt.Println(RunLengthDecode("12WB12W3B24WB"))
	// output: "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"
}
