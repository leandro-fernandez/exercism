package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	var hasLetter = strings.ContainsAny(strings.ToLower(remark), "abcdefghijklmnopqrstuvxwyz")
	var noLowerCase = strings.ToUpper(remark) == remark
	var isQuestion = strings.HasSuffix(remark, "?")

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case noLowerCase && hasLetter && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case noLowerCase && hasLetter:
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}
