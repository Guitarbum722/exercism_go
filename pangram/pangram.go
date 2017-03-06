// Package pangram provides utilities that will determine if input strings are a pangram.
package pangram

import (
	"strings"
)

const testVersion = 1

// IsPangram returns true if input contains all of the letters in alphabet.  The input must contain only ASCII characters.
func IsPangram(s string) bool {
	lowerLetters := make(map[rune]uint8)
	s = strings.ToLower(s)
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			lowerLetters[r] = 1
		}
	}
	return len(lowerLetters) == 26
}
