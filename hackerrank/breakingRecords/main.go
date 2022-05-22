package main

import "fmt"

// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {
	min := scores[0]
	max := scores[0]
	records := []int32{0, 0}
	for _, value := range scores {
		if value < min {
			records[1] += 1
			min = value
		} else if value > max {
			records[0] += 1
			max = value
		}
	}
	return records
}

func main() {
	scores := []int32{5, 2, 7, 6, 1, 5}
	answer := breakingRecords(scores)
	fmt.Println(answer)
}
