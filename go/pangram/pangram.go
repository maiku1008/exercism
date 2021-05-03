package pangram

import "unicode"

const alphabet = "abcdefghijklmnopqrstuvwxyz"


func IsPangram(sentence string) bool {
	letterSet := map[rune]bool{}
	for _, c := range sentence {
		if unicode.IsLetter(c) {
			letterSet[unicode.ToLower(c)] = true
		}
	}
	return len(letterSet) == len(alphabet)
}