// Package grains calculate grains
package grains

import (
	"errors"
	"math"
)

// Square calculates 2 in the power of val
func Square(val int) (uint64, error) {

	if val <= 0 || val > 64 {
		return 0, errors.New("invalid input")
	}

	var x, y float64 = float64(val) - 1, 2

	power := math.Pow(y, x)
	return uint64(power), nil
}

// Total returns the total value of grains in a chess board
func Total() uint64 {
	var sum uint64

	for i := 1; i <= 64; i++ {
		val, _ := Square(i)
		sum += val
	}

	return sum
}
