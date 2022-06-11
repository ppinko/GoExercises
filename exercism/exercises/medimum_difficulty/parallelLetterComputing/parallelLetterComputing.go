package main

import (
	"fmt"
)

// https://exercism.org/tracks/go/exercises/parallel-letter-frequency

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	half := len(l) / 2

	channel := make(chan FreqMap, 2)

	go mappers(l[0:half], channel)
	go mappers(l[half:], channel)

	total := FreqMap{}
	for i := 0; i < 2; i++ {
		m := <-channel
		for k, v := range m {
			total[k] += v
		}
	}

	return total
}

func mappers(slice []string, ch chan FreqMap) {
	occurrences := FreqMap{}
	for _, word := range slice {
		for _, char := range word {
			occurrences[char]++
		}
	}
	ch <- occurrences
}

func main() {
	words := []string{"banana", "apple", "pineapple", "watermelon"}
	word := ""
	for _, v := range words {
		word += v
	}

	fmt.Println()
	f1 := Frequency(word)
	f2 := ConcurrentFrequency(words)
	fmt.Println()
	fmt.Println(f1)
	fmt.Println(f2)
}
