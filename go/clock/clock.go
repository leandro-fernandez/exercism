// Package clock exposes a clock
package clock

import (
	"fmt"
)

// Clock represents a clock
type Clock struct {
	hour, minute int
}

// New creates a new clock
func New(hour, min int) Clock {
	var hourAdjusted, minAdjusted int

	minAdjusted = min % 60
	hour += min / 60

	if minAdjusted < 0 {
		minAdjusted += 60
		hour--
	}

	hourAdjusted = hour % 24

	if hourAdjusted < 0 {
		hourAdjusted += 24
	}

	return Clock{hourAdjusted, minAdjusted}
}

// Add adds time to an existing clock
func (clock Clock) Add(min int) Clock {
	return New(clock.hour, clock.minute+min)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}
