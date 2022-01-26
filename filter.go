package main

func filterWords(words []string, proposal string, result string) []string {
	filtered := []string{}

	for _, word := range words {
		rejected := false

		for _, policy := range filteringPolicies {
			if policy.reject(word, proposal, result) {
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
