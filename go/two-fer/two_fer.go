// Package twofer has the twofer func
package twofer

import (
	"fmt"
)

// ShareWith is a function that says 'one for {input}, one for me'
func ShareWith(s string) string {
	var name = "you"
	if s != "" {
		name = s
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
