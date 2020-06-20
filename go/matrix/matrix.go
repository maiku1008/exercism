package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// New takes an input string and uses it to initialize a matrix
func New(input string) (*Matrix, error) {
	split := strings.Split(input, "\n")

	m := &Matrix{}
	m.rows = make([][]int, len(split))
	var previous = -1 // keep track of previous iteration's slice length
	for i, v := range split {
		rowStr := strings.Fields(v)
		rowInt := make([]int, len(rowStr))
		for j, w := range rowStr {
			n, err := strconv.Atoi(w)
			if err != nil {
				return nil, err // found invalid character
			}
			rowInt[j] = n
		}
		m.rows[i] = rowInt

		switch {
		case previous == -1: // first iteration with no previous slice
			previous = len(m.rows[i])
		case previous != len(m.rows[i]):
			return nil, errors.New("found rows of different length")
		}
	}
	return m, nil
}

// Matrix is our matrix object
type Matrix struct {
	rows [][]int
}

// Rows returns a copy of the rows field of m
func (m *Matrix) Rows() [][]int {
	rows := make([][]int, len(m.rows))
	for i := range m.rows {
		rows[i] = make([]int, len(m.rows[i]))
		copy(rows[i], m.rows[i])
	}
	return rows
}

// Cols returns a copy of the cols field of m
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, len(m.rows[0]))
	for i := range cols {
		cols[i] = make([]int, len(m.rows))
		for j, row := range m.rows {
			cols[i][j] = row[i]
		}
	}
	return cols
}

// Set sets a value in the indicated row and column
func (m *Matrix) Set(row, col, val int) bool {
	switch {
	case row < 0 || row > len(m.rows)-1:
		return false
	case col < 0 || col > len(m.rows[0])-1:
		return false
	default:
		m.rows[row][col] = val
		return true
	}
}
