package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	length := 1 + (len(s) - n)
	result := make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = s[i : i+n]
	}
	return result
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First safely returns the first substring of s with length n.
func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
