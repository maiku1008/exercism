package scrabble

import "unicode"

// Score outputs the scrabble score of a word
func Score(word string) int {

	var result int
	for _, v := range word {
		switch unicode.ToLower(v) {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			result++
		case 'd', 'g':
			result += 2
		case 'b', 'c', 'm', 'p':
			result += 3
		case 'f', 'h', 'v', 'w', 'y':
			result += 4
		case 'k':
			result += 5
		case 'j', 'x':
			result += 8
		case 'q', 'z':
			result += 10
		}
	}

	return result
}
