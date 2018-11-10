// Package raindrops contains Convert
package raindrops

import (
	"strconv"
)

func _extractRaindrop(input, factor int, raindrop string) string {
	if input%factor == 0 {
		return raindrop
	}

	return ""
}

// Convert converts a number to its raindrops
func Convert(input int) string {
	raindrop := ""

	var raindrops = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for k, v := range raindrops {
		raindrop += _extractRaindrop(input, k, v)
	}

	if raindrop != "" {
		return raindrop
	}

	return strconv.Itoa(input)
}
