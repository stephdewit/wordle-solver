package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/stephdewit/wordle-solver/solver"
)

const length = 5

type Guess struct {
	Word   string `json:"w"`
	Result string `json:"r"`
}

type SuggestRequest struct {
	Guesses []Guess `json:"g"`
}

type SuggestResponse struct {
	Suggestions []string `json:"s"`
	Total       int      `json:"t"`
}

var cachedWords []solver.Word

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

	words := cachedWords
	for _, g := range req.Guesses {
		words = solver.FilterWords(words, g.Word, g.Result)
	}

	const maxSuggestions = 10
	suggestions := make([]string, 0, min(maxSuggestions, len(words)))
	for i, word := range words {
		if i >= maxSuggestions {
			break
		}
		suggestions = append(suggestions, word.Value)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(SuggestResponse{
		Suggestions: suggestions,
		Total:       len(words),
	})
}

func main() {
	var err error
	cachedWords, err = solver.LoadWords(length, false)
	if err != nil {
		log.Fatalf("failed to load word list: %v", err)
	}
	log.Printf("loaded %d words", len(cachedWords))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/suggestions", handleSuggest)

	addr := ":" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
