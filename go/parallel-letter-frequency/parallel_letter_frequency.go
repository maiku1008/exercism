package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency builds and returns a FreqMap by calling Frequency() concurrently
func ConcurrentFrequency(words []string) FreqMap {
	combined := FreqMap{}
	fmCh := make(chan FreqMap, 10)

	for _, word := range words {
		go func(s string) {
			fmCh <- Frequency(s)
		}(word)
	}

	for range words {
		for k, v := range <-fmCh {
			combined[k] += v
		}
	}

	return combined
}
