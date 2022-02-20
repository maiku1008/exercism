package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode compresses a string input by replacing consecutive characters
// with a count
func RunLengthEncode(input string) string {
	var encoded strings.Builder
	var count int

	for i, current := range input {
		count++
		// we are either at the end of the input,
		// or the current character differs from the next in the input
		if i == len(input)-1 || current != rune(input[i+1]) {
			switch {
			// If the character count is 0 or 1, just write the character
			case count <= 1:
				encoded.WriteRune(current)
			case count > 1:
				encoded.WriteString(strconv.Itoa(count))
				encoded.WriteRune(current)
			}
			count = 0
		}
	}
	return encoded.String()
}

// RunLengthDecode takes an encoded string and returns it to its original state
func RunLengthDecode(input string) string {
	var decoded strings.Builder
	multiplier := ""
	for _, current := range input {
		switch {
		// the character is a number so we add it to the multiplier string
		case unicode.IsNumber(current):
			multiplier += string(current)
		// when the character is not a number, then we write the correct number
		// of characters to the string
		case len(multiplier) > 0:
			s := string(current)
			m, _ := strconv.Atoi(multiplier)
			decoded.WriteString(strings.Repeat(s, m))
			multiplier = ""
		// the character is not a number, and also the length of the multiplier is 0
		// which means the character is a single letter
		default:
			decoded.WriteRune(current)
		}
	}
	return decoded.String()
}
