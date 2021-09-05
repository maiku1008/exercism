package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	wc := make(Frequency)
	input = strings.ToLower(input)
	words := strings.FieldsFunc(input, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})

	for _, word := range words {
		word = strings.Trim(word, "'")
		wc[word]++
	}
	return wc
}
