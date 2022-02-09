package main

type Word struct {
	value  string
	weight int
}

func word(value string, frequencies map[rune]int) Word {
	return Word{value: value, weight: getWeight(value, frequencies)}
}

func getWeight(word string, frequencies map[rune]int) int {
	w := 0
	knownLetters := make(map[rune]bool)

	for _, char := range []rune(word) {
		if knownLetters[char] {
			continue
		}

		knownLetters[char] = true
		w += frequencies[char]
	}

	return w
}
