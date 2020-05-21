package raindrops

import "strconv"

// Convert outputs a string which is different depending on num's factors
func Convert(num int) string {
	var output string
	if num%3 == 0 {
		output += "Pling"
	}

	if num%5 == 0 {
		output += "Plang"
	}

	if num%7 == 0 {
		output += "Plong"
	}

	if len(output) == 0 {
		output += strconv.Itoa(num)
	}

	return output
}
