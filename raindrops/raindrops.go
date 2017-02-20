// Package raindrops outputs a string based on a particular set of rules, or will simply output
// the input as a string if the criteria is not met.
package raindrops

import "strconv"

const testVersion = 2

// Convert outputs "Pling," "Plang" or Plong based on criteria.  Or, it will simply output the input value as a string if // the criteria is not met.
func Convert(num int) string {

	var out string

	if num%3 == 0 {
		out += "Pling"
	}
	if num%5 == 0 {
		out += "Plang"
	}
	if num%7 == 0 {
		out += "Plong"
	}
	if out == "" {
		out = strconv.Itoa(num)
	}
	return out
}
