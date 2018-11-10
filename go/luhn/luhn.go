// Package luhn has the luhn algorithm
package luhn

import (
	"strings"
)

func doubleDigit(digit int) int {
	double := digit * 2

	if double > 9 {
		double -= 9
	}

	return double
}

// Valid checks if an input is valid according to the luhn algorightm
func Valid(input string) bool {
	inputClean := strings.Replace(input, " ", "", -1)

	if len(inputClean) <= 1 {
		return false
	}

	sum := 0

	for index, char := range inputClean {
		charAsInt := int(char - '0')

		if index%2 == len(inputClean)%2 {
			sum += doubleDigit(charAsInt)
		} else {
			sum += charAsInt
		}
	}

	return sum%10 == 0
}
