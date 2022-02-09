package main

func filterWords(words []Word, proposal string, result string) []Word {
	filtered := []Word{}

	for _, word := range words {
		rejected := false

		for _, policy := range filteringPolicies {
			if policy.reject(word.value, proposal, result) {
				rejected = true
				break
			}
		}

		if rejected {
			continue
		}

		filtered = append(filtered, word)
	}

	return filtered
}
