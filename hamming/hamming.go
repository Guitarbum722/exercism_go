/*
Package hamming provides a utility that determines the hamming distance of two strands of DNA.
*/
package hamming

import "errors"

const testVersion = 5

// Distance accepts two strings representing strands of DNA, and returns the hamming distance of 0 or more.  Note that 0 will also be returned if err is not nil.
func Distance(a, b string) (int, error) {
	// if a == "" || b == "" {
	// 	return 0, errors.New("Two input strings must be provided")
	// }
	if len(a) != len(b) {
		return -1, errors.New("Inputs must be of equal length")
	}

	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
