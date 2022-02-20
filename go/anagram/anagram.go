package anagram

import (
	"sort"
	"strings"
)

// Detect determines if a list of candidate words are anagrams of a subject word
func Detect(subject string, candidates []string) []string {
	s := strings.Split(strings.ToLower(subject), "")
	sort.Strings(s)

	found := []string{}
	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}
		if len(subject) != len(candidate) {
			continue
		}
		c := strings.Split(strings.ToLower(candidate), "")
		sort.Strings(c)
		if same(s, c) {
			found = append(found, candidate)
		}
	}
	return found
}

// same determines if two slices of string have the same length
// and contain the same elements
func same(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
