package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	transformed := map[string]int{}
	for score, letters := range in {
		for i := range letters {
			transformed[strings.ToLower(letters[i])] = score
		}
	}
	return transformed
}
