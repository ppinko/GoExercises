package main

import (
	"fmt"
	"testing" // package for testing
)

func TestAdd(t *testing.T) { // pointer for testing object for logging
	x, y, expectedResult := 1, 2, 4
	if Add(x, y) != expectedResult {
		// log message and fails now
		t.Fatalf(`Result of sum %d and %d should equal %d`, 1, 2, 3)
	}
}

func TestAddTable(t *testing.T) {
	// table-driven style
	var tests = []struct {
		x, y      int
		expResult int
	}{
		{1, 1, 2},
		{2, -2, 0},
		{-1, 5, 4},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.x, tt.y)
		t.Run(testname, func(t *testing.T) { // name the subtest
			ans := Add(tt.x, tt.y)
			if ans != tt.expResult {
				t.Errorf("got %d, want %d", ans, tt.expResult)
			}
		})
	}
}
