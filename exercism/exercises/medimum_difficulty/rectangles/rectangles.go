package main

// https://exercism.org/tracks/go/exercises/rectangles

// constants
const (
	corner rune = rune('+')
	h      rune = rune('-')
	v      rune = rune('|')
)

type Point struct {
	row int
	col int
}

func ValidRectangular(c1, c2 Point, diag [][]rune) bool {
	// validate horizontal sides
	for i := c1.col + 1; i < c2.col; i++ {
		if (diag[c1.row][i] != h && diag[c1.row][i] != corner) ||
			(diag[c2.row][i] != h && diag[c2.row][i] != corner) {
			return false
		}
	}

	// validate vertical sides
	for i := c1.row + 1; i < c2.row; i++ {
		if (diag[i][c1.col] != v && diag[i][c1.col] != corner) ||
			(diag[i][c2.col] != v && diag[i][c2.col] != corner) {
			return false
		}
	}

	return true
}

func Count(diagram []string) int {
	// convert to diagram of runes
	diag := make([][]rune, len(diagram))
	for i, str := range diagram {
		diag[i] = []rune(str)
	}

	numOfRectangulars := 0
	for i := 0; i < len(diag); i++ {
		for j := 0; j < len(diag[i]); j++ {
			if diag[i][j] == corner {
				for k := j + 1; k < len(diag[i]); k++ {
					if diag[i][k] == corner {
						for l := i + 1; l < len(diag); l++ {
							if diag[l][j] == corner && diag[l][k] == corner {
								if ValidRectangular(Point{i, j}, Point{l, k}, diag) {
									numOfRectangulars++
								}
							}
						}
					}
				}
			}
		}
	}

	return numOfRectangulars
}

func main() {
}
