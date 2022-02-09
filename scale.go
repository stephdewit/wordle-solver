package main

func getFrequencies(words []string) map[rune]int {
	freq := make(map[rune]int)

	for _, word := range words {
		for _, char := range []rune(word) {
			freq[char] = freq[char] + 1
		}
	}

	return freq
}
