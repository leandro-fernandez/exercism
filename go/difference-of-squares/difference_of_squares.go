// Package diffsquares has a func to calculate the difference of squares
package diffsquares

// SquareOfSums calculates the square of the sums of an integer number
func SquareOfSums(input int) int {
	sum := 0

	for i := 1; i <= input; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares calculates the sum of the squares of an integer number
func SumOfSquares(input int) int {
	sum := 0

	for i := 1; i <= input; i++ {
		sum += i * i
	}

	return sum
}

// Difference calculates the difference between the square of the sums and the sum of squares of an integer number
func Difference(input int) int {
	return SquareOfSums(input) - SumOfSquares(input)
}
