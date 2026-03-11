package solver

import "sort"

// BestProbes returns words from allWords that best reveal unknown letters
// among the remaining candidates, ranked by summed letter frequency.
func BestProbes(candidates []Word, allWords []Word, guessedWords []string) []Word {
	strs := make([]string, len(candidates))
	for i, w := range candidates {
		strs[i] = w.Value
	}
	freq := getFrequencies(strs)

	tested := map[rune]bool{}
	for _, w := range guessedWords {
		for _, ch := range w {
			tested[ch] = true
		}
	}

	type scored struct {
		word  Word
		score int
	}
	var probes []scored
	for _, w := range allWords {
		score := 0
		seen := map[rune]bool{}
		for _, ch := range []rune(w.Value) {
			if !tested[ch] && !seen[ch] {
				score += freq[ch]
				seen[ch] = true
			}
		}
		if score > 0 {
			probes = append(probes, scored{w, score})
		}
	}

	sort.Slice(probes, func(i, j int) bool {
		return probes[i].score > probes[j].score
	})

	const max = 10
	result := make([]Word, 0, min(max, len(probes)))
	for i, p := range probes {
		if i >= max {
			break
		}
		result = append(result, p.word)
	}
	return result
}
