package main

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

type numberBoxContaining struct {
	Value int
}

func (n numberBoxContaining) Number() int {
	return n.Value
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type AnotherFancyNumber struct {
	n string
}

func (i AnotherFancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		val, _ := strconv.Atoi(fnb.Value())
		return val
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(num))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	default:
		return "Return to sender"
	}
}

func main() {
	fmt.Println(DescribeNumber(-12.345))
	// Output: This is the number -12.3

	fmt.Println(DescribeNumberBox(numberBoxContaining{12}))
	// Output: This is a box containing the number 12.0

	fmt.Println(ExtractFancyNumber(FancyNumber{"10"}))
	// Output: 10
	fmt.Println(ExtractFancyNumber(AnotherFancyNumber{"4"}))
	// Output: 0

	fmt.Println(DescribeFancyNumberBox(FancyNumber{"10"}))
	// Output: This is a fancy box containing the number 10.0

	fmt.Println(DescribeFancyNumberBox(AnotherFancyNumber{"4"}))
	// Output: This is a fancy box containing the number 0.0

	// fmt.Println(DescribeAnything(numberBoxContaining{12.345}))
	// Output: This is a box containing the number 12.3

	fmt.Println(DescribeAnything("some string"))
	// Output: Return to sender
}
