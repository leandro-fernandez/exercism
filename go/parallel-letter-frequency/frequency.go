package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(input []string) FreqMap {
	freqChan := make(chan FreqMap)

	for _, phrase := range input {
		go func(phrase string) {
			freqChan <- Frequency(phrase)
		}(phrase)
	}

	freqMap := make(FreqMap)

	for range input {
		for letter, freq := range <-freqChan {
			freqMap[letter] += freq
		}
	}

	return freqMap
}