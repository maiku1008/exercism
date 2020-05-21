package accumulate

// Accumulate runs a function convFn on a collection and returns another collection with the result
func Accumulate(collection []string, convFn func(string) string) []string {
	result := make([]string, len(collection))
	for i := range collection {
		result[i] = convFn(collection[i])
	}

	return result
}
