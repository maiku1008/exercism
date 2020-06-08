package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const max = 26 * 26 * 10 * 10 * 10

var used = map[string]bool{}

// Robot defines a robot with a name
type Robot struct {
	name string
}

// Name generates a new name for the robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	// Exhausted unique possibilities
	if len(used) >= max {
		return "", errors.New("used up all possible names")
	}

	r.name = generate()
	for used[r.name] {
		r.name = generate()
	}
	used[r.name] = true

	return r.name, nil
}

// Reset wipes the name from the robot's memory
func (r *Robot) Reset() {
	r.name = ""
}

// Helper function to generate a random robot name
// in the form AA###
func generate() string {
	r1 := rand.Intn(26) + 'A'
	r2 := rand.Intn(26) + 'A'
	num := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
