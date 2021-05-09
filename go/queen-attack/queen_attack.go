package queenattack

import (
	"errors"
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

	// Only set the initial position for black
	b, err := newQueen(black)
	if err != nil {
		return false, err
	}

	// If any of the square of white match with black's position,
	// it means the queens can attack each other
	if w.board&b.board != 0 {
		return true, nil
	}
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
	board uint64
	rank  int
	file  int
}

func (b *queen) setPosition(file, rank int) {
	bit := file + (rank * width)
	b.setBit(bit)
}

func (b *queen) setBit(pos int) {
	b.board |= (1 << pos)
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
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			if b.file-b.rank == x-y {
				b.setPosition(x, y)
			}
			if b.file+b.rank == x+y {
				b.setPosition(x, y)
			}
		}
	}
}
