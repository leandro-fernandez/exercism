// Package gigasecond provides a func to calculate when a person is a gigasecond old
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond receives a person's time of birth and returns when that person is a gigasecond old
func AddGigasecond(t time.Time) time.Time {
	var gigasecond = 1000000000 * time.Second
	return t.Add(gigasecond)
}
