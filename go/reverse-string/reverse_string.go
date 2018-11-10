// Package reverse has a func to reverse strings
package reverse

// String returns `input` reversed
func String(input string) string {

	reversed := ""

	for _, char := range input {
		reversed = string(char) + reversed
	}

	return reversed
}
