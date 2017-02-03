/*
Package clock provides a general purpose clock utility
*/

package clock

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}
}

func (Clock) String() string {
}

func (Clock) Add(minutes int) Clock {
}
