package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber returns a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox returns a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber returns the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fn, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}
	res, _ := strconv.Atoi(fn.Value())
	return res
}

// DescribeFancyNumberBox returns a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fn := fnb.(type) {
	case FancyNumber:
		n, _ := strconv.Atoi(fn.Value())
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(n))
	default:
		return "This is a fancy box containing the number 0.0"
	}
}

// DescribeAnything returns a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch n := i.(type) {
	case int:
		return DescribeNumber(float64(n))
	case float64:
		return DescribeNumber(n)
	case NumberBox:
		return DescribeNumberBox(n)
	case FancyNumberBox:
		return DescribeFancyNumberBox(n)
	default:
		return "Return to sender"
	}
}
