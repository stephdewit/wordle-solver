package main

import (
	"encoding/json"
	"net/http"

	"github.com/stephdewit/wordle-solver/solver"
)

func handleSuggest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SuggestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	candidates := cachedWords
	guessedWords := make([]string, len(req.Guesses))
	for i, g := range req.Guesses {
		candidates = solver.FilterWords(candidates, g.Word, g.Result)
		guessedWords[i] = g.Word
	}

	const max = 10
	suggestions := make([]string, 0, min(max, len(candidates)))
	for i, word := range candidates {
		if i >= max {
			break
		}
		suggestions = append(suggestions, word.Value)
	}

	probeWords := solver.BestProbes(candidates, cachedWords, guessedWords)
	probes := make([]string, len(probeWords))
	for i, word := range probeWords {
		probes[i] = word.Value
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SuggestResponse{
		Suggestions: suggestions,
		Probes:      probes,
		Total:       len(candidates),
	})
}
