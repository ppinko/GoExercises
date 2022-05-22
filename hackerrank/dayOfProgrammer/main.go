package main

import (
	"fmt"
	"strconv"
)

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	firstEightMonths := []int8{31, 28, 31, 30, 31, 30, 31, 31}
	var sum int
	for _, value := range firstEightMonths {
		sum += int(value)
	}
	if year >= 1919 {
		if (year%400 == 0) || (year%100 != 0 && year%4 == 0) {
			sum += 1
		}
	} else if year == 1918 {
		sum -= 13

	} else {
		if year%4 == 0 {
			sum += 1
		}
	}
	remainder := 256 - sum
	var date string
	if remainder < 10 {
		date += "0"
	}
	date += strconv.Itoa(remainder) + ".09." + strconv.Itoa(int(year))
	return date
}

func main() {
	fmt.Println(dayOfProgrammer(1918))
}
