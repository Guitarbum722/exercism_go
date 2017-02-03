/*
Package leap provides utilities that determine whether or not a year is a leap year.
*/
package leap

const testVersion = 3

// IsLeapYear returns true if the input year is a leap year
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}
