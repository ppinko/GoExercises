package main

import (
	"errors"
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	max := int(byte('Z') - byte('A'))
	i := int(char - byte('A'))
	if i < 0 || i > max {
		return "", errors.New("wrong input")
	}
	width := 2*i + 1
	mid := i
	diamond := make([]string, width)
	val := 'A'
	for j := 0; j <= mid; j++ {
		line := strings.Repeat(" ", int(width))

		line = line[:mid-j] + string(val) + line[mid-j+1:]
		line = line[:mid+j] + string(val) + line[mid+j+1:]
		diamond[j] = line
		diamond[width-j-1] = line

		val++
	}
	ret := strings.Join(diamond, "\n")
	return ret, nil
}

func main() {
	fmt.Println(Gen(byte('F')))
}
