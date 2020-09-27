package gigasecond

import "time"

// AddGigasecond adds gigasecond to a time t
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
