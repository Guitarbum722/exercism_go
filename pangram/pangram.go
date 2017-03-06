// Package pangram provides utilities that will determine if input strings are a pangram.
package pangram

import (
	"strings"
)

const testVersion = 1

// IsPangram returns true if input contains all of the letters in alphabet.  The input must contain only ASCII characters.
func IsPangram(input string) bool {
	var letterList = make(map[rune]int)

	if input == "" {
		return false
	}
	for r := 'a'; r <= 'z'; r++ {
		letterList[r] = 0
	}
	input = strings.ToLower(input)
	for _, v := range input {
		if v < 'a' || v > 'z' {
			continue
		}
		letterList[v] = 1
	}
	for k := range letterList {
		if letterList[k] < 1 {
			return false
		}
	}
	return true
}
