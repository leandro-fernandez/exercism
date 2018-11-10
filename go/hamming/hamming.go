// Package hamming contains Distance
package hamming

import (
	"errors"
)

// Distance calculates the humming distance between two dna strands of the same size
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA strands are not of equal size")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
