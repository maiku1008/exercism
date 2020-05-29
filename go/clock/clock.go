package clock

import "fmt"

// New returns a new Clock
func New(hour, minute int) Clock {
	var minutes = (hour * 60) + minute
	minutes %= 24 * 60
	if minutes < 0 {
		minutes += 24 * 60
	}

	return Clock{minutes: minutes}
}

// Clock represents a clock object that stores minutes
type Clock struct {
	minutes int
}

// Add adds minutes to a clock and returns a new clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes + minutes)
}

// Subtract subtracts minutes from a clock object and returns a new clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes - minutes)
}

// String returns a string representation of the clock
// with hours 0-23 minutes 0-59
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
