package letter

import (
	"sync"
)

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
	var wg sync.WaitGroup
	var lock sync.RWMutex

	for _, word := range words {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for k, v := range Frequency(s) {
				lock.Lock()
				combined[k] += v
				lock.Unlock()
			}
		}(word)
	}
	wg.Wait()
	return combined
}
