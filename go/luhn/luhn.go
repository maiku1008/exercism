package luhn

import (
	"strconv"
	"strings"
)

// Valid checks whether an input string is validated by Luhn's algorithm
func Valid(input string) bool {

	// Remove whitespace
	input = strings.ReplaceAll(input, " ", "")

	// Check length
	if len(input) < 2 {
		return false
	}

	var (
		even = len(input)%2 == 0
		sum  int
	)

	for _, v := range input {
		// Convert string value to int.
		// If the value is not a number, return false.
		digit, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}

		// Determine if we need to double our digit according to even bool
		if even {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		even = !even
		sum += digit
	}

	return sum%10 == 0
}
