// Package acronym provids utilities for converting phrases or names to their respective acronyms.
package acronym

import "strings"
import "regexp"

const testVersion = 2

// Abbreviate accepts a string.  Input full is converted to its acronym and is returned as res.
func Abbreviate(full string) string {
	var res []rune
	full = strings.Replace(full, "-", " ", -1)
	words := strings.Split(full, " ")

	// if first contains a colon then it is assumed that the acronym is already defined and it will be returned.
	if ok, _ := regexp.MatchString(":", words[0]); ok {
		return strings.TrimSuffix(words[0], ":")
	}
	for _, word := range words {
		res = append(res, rune(word[0]))
		for _, v := range word[1:] {
			if v >= 'A' && v <= 'Z' {
				res = append(res, rune(v))
			}
		}
	}
	return strings.ToUpper(string(res))
}
