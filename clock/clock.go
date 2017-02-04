/*
Package clock provides a general purpose clock utility
*/

package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

type Clock struct {
	hour, minute int
}

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
	fmt.Println(hour, minute)
	return Clock{hour, minute}
}

func (c Clock) String() string {
	out := fmt.Sprintf("%02d:%02d", c.hour, c.minute)
	return out
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c
}
