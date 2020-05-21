package isogram

import "unicode"

// IsIsogram returns true if a word is an isogram
func IsIsogram(word string) bool {
	set := make(map[rune]bool)
	for _, v := range word {
		if v == '-' || v == ' ' {
			continue
		}
		v = unicode.ToLower(v)
		if set[v] {
			return false
		}
		set[v] = true
	}
	return true
}
