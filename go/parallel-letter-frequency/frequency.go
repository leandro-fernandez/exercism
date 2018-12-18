package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func FrequencyWithMap(s string, m FreqMap, done chan bool) FreqMap {
	for _, r := range s {
		m[r]++
	}

	done <- true

	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}
	done := make(chan bool)

	for index, _ := range s {
		go FrequencyWithMap(s[index], m, done)
	}

	for i := 0; i < len(s); i++ {
		<-done
	}

	return m
}