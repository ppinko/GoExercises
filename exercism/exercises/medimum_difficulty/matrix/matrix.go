package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// https://exercism.org/tracks/go/exercises/matrix

// Define the Matrix type here.

type Matrix [][]int

func New(s string) (*Matrix, error) {
	var matrix Matrix

	if len(s) == 0 {
		return &matrix, errors.New("empty string")
	}

	cols := 0
	rows := strings.Split(s, "\n")

	for i, v := range rows {
		line := strings.Trim(v, " \n\t\r")
		nums := strings.Split(line, " ")
		var row []int
		for _, el := range nums {
			num, err := strconv.Atoi(el)
			if err != nil {
				return &matrix, errors.New("wrong input")
			}
			row = append(row, num)
		}
		if i == 0 {
			cols = len(row)
		} else if len(row) != cols {
			return &matrix, errors.New("wrong input")
		}

		matrix = append(matrix, row)
	}
	return &matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	var cols [][]int

	for i := 0; i < len((*m)[0]); i++ {
		var col []int
		for j := 0; j < len(*m); j++ {
			col = append(col, (*m)[j][i])
		}
		cols = append(cols, col)
	}
	return cols
}

func (m *Matrix) Rows() [][]int {
	arr := make([][]int, len(*m))
	for i := 0; i < len(*m); i++ {
		r := make([]int, len((*m)[0]))
		copy(r, (*m)[i])
		arr[i] = r
	}
	return arr
}

func (m *Matrix) Set(row, col, val int) bool {
	// do not allow negative indices
	if row < 0 || col < 0 {
		return false
	}
	// check index out of range
	if len(*m) <= row || len((*m)[0]) <= col {
		return false
	} else {
		(*m)[row][col] = val
		return true
	}
}
func main() {
	fmt.Println()

	test1 := "1 2\n10 20"

	matrix, _ := New(test1)
	fmt.Println(*matrix)
}
