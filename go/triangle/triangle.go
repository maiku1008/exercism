package triangle

import "math"

// Kind is the kind of triangle
type Kind int

const (
	// NaT indicates Not a triangle
	NaT = iota
	// Equ indicates an equilateral triangle
	Equ
	// Iso indicates an isosceles triangle
	Iso
	// Sca indicates a scalene triangle
	Sca
)

// KindFromSides returns the kind of triangle according to measurements of its 3 sides
func KindFromSides(a, b, c float64) Kind {

	if Valid(a, b, c) {
		switch {
		case a == b && a == c:
			return Equ
		case (a == b) || (b == c) || (c == a):
			return Iso
		case a != b && a != c:
			return Sca
		}
	}
	return NaT
}

// Valid checks whether 3 segments make a valid triangle
func Valid(a, b, c float64) bool {
	return (a > 0 && b > 0 && c > 0) &&
		((a+b >= c) && (a+c >= b) && (b+c >= a)) &&
		(!math.IsInf(a, 0) && !math.IsInf(b, 0) && !math.IsInf(c, 0))
}
