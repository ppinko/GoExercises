package main

import (
	"fmt"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	app := map[rune]string{
		'\u2757':     "recommendation",
		'\U0001f50d': "search",
		'\u2600':     "weather",
	}
	for _, char := range log {
		if name, ok := app[char]; ok {
			return name
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for index, char := range runes {
		if char == oldRune {
			runes[index] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runes := []rune(log)
	return len(runes) <= limit
}

func main() {
	fmt.Println(Application("❗ recommended search product 🔍"))

	log := "please replace '👎' with '👍'"

	fmt.Println(Replace(log, '👎', '👍'))
	// Output: please replace '👍' with '👍'"

	fmt.Println(WithinLimit("hello❗", 6))
}
