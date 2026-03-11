package solver

type Word struct {
	Value  string
	weight int
}

func newWord(value string, frequencies map[rune]int) Word {
	return Word{Value: value, weight: getWeight(value, frequencies)}
}

func getWeight(w string, frequencies map[rune]int) int {
	weight := 0
	knownLetters := make(map[rune]bool)

	for _, char := range []rune(w) {
		if knownLetters[char] {
			continue
		}

		knownLetters[char] = true
		weight += frequencies[char]
	}

	return weight
}
