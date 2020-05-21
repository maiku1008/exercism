package hamming

import (
	"fmt"
)

// Distance calculates the hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, fmt.Errorf("strings should be of same length")
	}

	var distance int
	for i := 0; i < len(ar); i++ {
		if ar[i] != br[i] {
			distance++
		}
	}
	return distance, nil
}
