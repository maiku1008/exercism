package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
	"sync"
)

const (
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	max = 26*26*10*10*10
)

var used = map[string]bool{}

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {

	if r.name != "" {
		return r.name, nil
	}

	// Exhausted unique possibilities
	if len(used) >= max {
		return "", errors.New("used up all possible names")
	}

	var mu sync.RWMutex
	resCh := make(chan string, 10)

	var ok bool
	for !ok {

		// fmt.Println("looped: ", n) // Debugging
		go func() {
			n := generate()
			mu.Lock()
			if !used[n] {
				used[n] = true
				resCh <- n
			}
			mu.Unlock()
		}()
	}

	for n := range resCh {
		mu.Lock()
		mu.Unlock()
	}


	r.name = n
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

// Helper function to generate a random robot name
// in the form AA###
func generate() string {
	rand.Seed(time.Now().UnixNano())

	var name string
	for i := 0; i < 2; i++ {
		name += string(chars[rand.Intn(len(chars))])
	}
	for i := 0; i < 3; i++ {
		name += strconv.Itoa(rand.Intn(9))
	}
	return name
}
