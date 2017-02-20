// Package raindrops outputs a string based on a particular set of rules, or will simply output
// the input as a string if the criteria is not met.
package raindrops

import "fmt"

const testVersion = 2

// Convert outputs "Pling," "Plang" or Plong based on criteria.  Or, it will simply output the input value as a string if // the criteria is not met.
func Convert(num int) string {

	var out string
	var div3 = num%3 == 0
	var div5 = num%5 == 0
	var div7 = num%7 == 0

	if !(div3 || div5 || div7) {
		out = fmt.Sprintf("%d", num)
	}
	if div3 {
		out = fmt.Sprintf("%s%s", out, "Pling")
	}
	if div5 {
		out = fmt.Sprintf("%s%s", out, "Plang")
	}
	if div7 {
		out = fmt.Sprintf("%s%s", out, "Plong")
	}
	return out
}
