/*
Package clock provides a general purpose clock utility
*/

package clock

import "fmt"

const testVersion = 4

type Clock struct {
	hour, minute int
}

// Creates a new clock and rolls over minutes and hours
func New(hour, minute int) Clock {
	for minute < 0 {
		minute = 60 + minute
		hour--
	}
	if minute > 59 {
		hour += minute / 60
		minute = minute % 60
	}
	if hour >= 24 {
		hour = hour % 24
	}

	for hour < 0 {
		hour = 24 + hour
	}
	return Clock{hour, minute}
}

// Stringify the clock
func (c Clock) String() string {
	out := fmt.Sprintf("%02d:%02d", c.hour, c.minute)
	return out
}

// Add input even if negative and roll over hours and minutes appropriately
func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	for c.minute < 0 {
		c.minute = 60 + c.minute
		c.hour--
	}
	if c.minute > 59 {
		c.hour += c.minute / 60
		c.minute = c.minute % 60
	}
	if c.hour >= 24 {
		c.hour = c.hour % 24
	}

	for c.hour < 0 {
		c.hour = 24 + c.hour
	}
	return c
}
