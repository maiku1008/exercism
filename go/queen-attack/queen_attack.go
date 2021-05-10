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
	setAttackSquares := true
	w, err := newBoard(white, setAttackSquares)
	if err != nil {
		return false, err
	}
	b, err := newBoard(black, !setAttackSquares)
	if err != nil {
		return false, err
	}
	if w&b != 0 {
		return true, nil
	}
	return false, nil
}

func newBoard(position string, setAttackSquares bool) (bitboard, error) {
	if !(len(position) == 2) {
		return 0, errors.New("invalid square")
	}
	file, rank := position[0], position[1]
	if file < 'a' || file > 'h' {
		return 0, errors.New("square out of bounds")
	}
	if rank < '1' || rank > '8' {
		return 0, errors.New("square out of bounds")
	}
	x, y := int(file-'a'), int(rank-'1')

	var q bitboard
	q.setBit(x + (y * width))

	if setAttackSquares {
		q.setFile(x, y)
		q.setRank(x, y)
		q.setDiagonal(x, y)
	}
	return q, nil
}

type bitboard uint64

func (b *bitboard) setBit(pos int) {
	*b |= (1 << pos)
}

func (b *bitboard) setFile(file, rank int) {
	for i := 0; i < width; i++ {
		b.setBit(i + (rank * width))
	}
}

func (b *bitboard) setRank(file, rank int) {
	for i := 0; i < width; i++ {
		b.setBit(file + (i * width))
	}
}

func (b *bitboard) setDiagonal(file, rank int) {
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			if file-rank == x-y {
				b.setBit(x + (y * width))
			}
			if file+rank == x+y {
				b.setBit(x + (y * width))
			}
		}
	}
}
