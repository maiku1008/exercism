package thefarm

import (
	"errors"
	"fmt"
)

var (
	errNegativeFodder = errors.New("negative fodder")
	errDivisionByZero = errors.New("division by zero")
)

// SillyNephewError is a custom error that makes fun of the silly nephew
// and also returns the number of cows.
type SillyNephewError struct {
	Cows int
}

func (n *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", n.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(wf WeightFodder, cows int) (float64, error) {
	fodder, err := wf.FodderAmount()
	switch {
	case fodder > 0 && err == ErrScaleMalfunction:
		fodder *= 2
	case fodder < 0 && (err == ErrScaleMalfunction || err == nil):
		return 0, errNegativeFodder
	case cows == 0:
		return 0, errDivisionByZero
	case cows < 0:
		return 0, &SillyNephewError{Cows: cows}
	case err != nil:
		return 0, err
	}
	return fodder / float64(cows), nil
}
