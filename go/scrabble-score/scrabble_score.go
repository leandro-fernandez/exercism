// Package scrabble contains Score to calculate scrabble scores
package scrabble

import (
	"strings"
)

// Score calculates the scrabble score for a given input
func Score(input string) int {

	var lettersToPoints = map[rune]int{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'l': 1, 'n': 1, 'r': 1, 's': 1, 't': 1,
		'd': 2, 'g': 2,
		'b': 3, 'c': 3, 'm': 3, 'p': 3,
		'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
		'k': 5,
		'j': 8, 'x': 8,
		'q': 10, 'z': 10,
	}

	points := 0

	inputLower := strings.ToLower(input)

	for _, char := range inputLower {
		points += lettersToPoints[char]
	}

	return points
}
