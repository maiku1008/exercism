package luhn

import (
	"strconv"
	"strings"
)

// Valid checks whether an input string is validated by Luhn's algorithm
func Valid(input string) bool {

	// Remove whitespace
	input = strings.Replace(input, " ", "", -1)

	// Check length
	if len(input) < 2 {
		return false
	}

	// Convert the string into digits.
	// return false if the character is not a number.
	var digits = make([]int, len(input))
	for i, v := range input {
		d, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		digits[i] = d
	}

	// Double every second digit starting from the right
	i := len(digits) - 2
	for i >= 0 {
		double := digits[i] * 2
		if double > 9 {
			double -= 9
		}
		digits[i] = double
		i -= 2
	}

	// Sum all digits
	var sum int
	for _, v := range digits {
		sum += v
	}

	// Validate
	if sum%10 != 0 {
		return false
	}

	return true
}
