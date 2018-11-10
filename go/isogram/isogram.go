// Package isogram contains a func to determine if a word is an isogram
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns `true` if the input is a isogram word, otherwise returns false
func IsIsogram(word string) bool {

	letterExists := map[rune]bool{}

	for _, char := range strings.ToLower(word) {
		if letterExists[char] {
			return false
		} else if unicode.IsLetter(char) {
			letterExists[char] = true
		}
	}

	return true
}
