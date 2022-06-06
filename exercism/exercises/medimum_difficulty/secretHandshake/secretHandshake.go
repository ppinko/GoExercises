package main

import "fmt"

// https://exercism.org/tracks/go/exercises/secret-handshake

func Handshake(code uint) []string {
	actions := map[uint]string{1: "wink",
		2: "double blink",
		4: "close your eyes",
		8: "jump"}
	keys := []uint{1, 2, 4, 8}

	const reverse uint = 16

	var handshake []string
	for _, k := range keys {
		if code&k == k {
			handshake = append(handshake, actions[k])
		}
	}
	// check if reverse is required
	if code&reverse == reverse {
		// reverse the
		l := len(handshake)
		for i := 0; i < l/2; i++ {
			// swap values
			handshake[i], handshake[l-i-1] = handshake[l-i-1], handshake[i]
		}
	}
	return handshake
}

func main() {
	fmt.Println(Handshake(19))
}
