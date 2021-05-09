package queenattack

import (
	"errors"
	"fmt"
)

const width = 8

// CanQueenAttack determines whether two queens on a chessboard
// are able to attack each other
func CanQueenAttack(white, black string) (bool, error) {
	if white == black {
		return false, errors.New("positions for white and black are the same")
	}
	w, err := newQueen(white)
	if err != nil {
		return false, err
	}
	// Fill the board with possible attack positions from white
	w.setFile()
	w.setRank()
	w.setDiagonal()

	// only set initial position for black
	b, err := newQueen(black)
	if err != nil {
		return false, err
	}

	// If any of the attack positions of white match with black's position,
	// queens can attack each other
	if w.position&b.position != 0 {
		return true, nil
	}
	// Print the boards for debugging purposes
	w.print()
	fmt.Println()
	b.print()
	// n, _ := newQueen("a1")
	// n.print()
	return false, nil
}

func newQueen(position string) (*queen, error) {
	if !(len(position) == 2) {
		return nil, errors.New("invalid square")
	}
	file, rank := position[0], position[1]
	if file < 'a' || file > 'h' {
		return nil, errors.New("square out of bounds")
	}
	if rank < '1' || rank > '8' {
		return nil, errors.New("square out of bounds")
	}
	q := &queen{
		file: int(file - 'a'),
		rank: int(rank - '1'),
	}
	q.setPosition(q.file, q.rank)

	return q, nil
}

type queen struct {
	position uint64
	rank     int
	file     int
}

func (b *queen) setPosition(file, rank int) {
	bit := file + (rank * width)
	b.setBit(bit)
}

func (b *queen) setBit(pos int) {
	b.position |= (1 << pos)
}

func (b *queen) setFile() {
	for i := 0; i < width; i++ {
		b.setBit(i + (b.rank * width))
	}
}

func (b *queen) setRank() {
	for i := 0; i < width; i++ {
		b.setBit(b.file + (i * width))
	}
}

func (b *queen) setDiagonal() {
	diagonal := b.rank - b.file + 7
	antidiagonal := b.rank + b.file
	for i := 0; i < width; i++ {
		for j := 0; j < width; j++ {
			diag := i - j + 7
			antid := j + i
			if diag == diagonal || antid == antidiagonal {
				b.setPosition(i, j)
			}
		}
	}
}

func (b *queen) print() {
	board := fmt.Sprintf("%064b", b.position)
	result := string(board[0])
	for i := 1; i < width*width; i++ {
		if i%width == 0 {
			result += "\n"
		}
		result += string(board[i])
	}
	fmt.Println(result)
	fmt.Println()
}
