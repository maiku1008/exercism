package queenattack

import (
	"errors"
)

// CanQueenAttack determines whether two queens on a chessboard
// are able to attack each other
func CanQueenAttack(white, black string) (bool, error) {
	if !valid(white) || !valid(black) || white == black {
		return false, errors.New("invalid input")
	}
	w := square{file: white[0], rank: white[1]}
	b := square{file: black[0], rank: black[1]}

	if sameFile(w, b) {
		return true, nil
	}
	if sameRank(w, b) {
		return true, nil
	}
	if sameDiagonal(w, b) {
		return true, nil
	}
	return false, nil
}

type square struct {
	file byte
	rank byte
}

func valid(pos string) bool {
	return len(pos) == 2 &&
		pos[0] >= 'a' &&
		pos[0] <= 'h' &&
		pos[1] >= '1' &&
		pos[1] <= '8'
}

func sameFile(w, b square) bool {
	return w.file == b.file
}

func sameRank(w, b square) bool {
	return w.rank == b.rank
}

func sameDiagonal(w, b square) bool {
	return w.file-b.file == w.rank-b.rank || w.file-b.file == b.rank-w.rank
}
