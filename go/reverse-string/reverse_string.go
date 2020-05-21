package reverse

// Reverse returns a reversed word
func Reverse(word string) string {
	var result = []rune(word)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return string(result)
}
