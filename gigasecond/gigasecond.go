/*
Package gigasecond provides utilities to determine when someone has lived for 1 gigasecond.
*/
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// AddGigasecond Adds a gigisecond to the input time
func AddGigasecond(tt time.Time) time.Time {
	gsec := time.Duration(1000000000)
	return tt.Add(time.Second * gsec)
}
