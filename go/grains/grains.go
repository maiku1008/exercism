package grains

import "fmt"

// Square returns the number of grains on the chosen square
// Returns an error if the square is out of accepted range
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, fmt.Errorf("square %d out of range 1-64", square)
	}

	return 1 << (square - 1), nil
}

// Total sums all the grains on the board
func Total() uint64 {
	return 1<<64 - 1
}
