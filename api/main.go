package main

import (
	"encoding/json"
	"flag"
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

func cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

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
	json.NewEncoder(w).Encode(SuggestResponse{
		Suggestions: suggestions,
		Total:       len(words),
	})
}

func main() {
	enableCORS := flag.Bool("cors", false, "enable CORS middleware")
	flag.Parse()

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

	handler := http.HandlerFunc(handleSuggest)
	if *enableCORS {
		handler = cors(handleSuggest)
		log.Println("CORS enabled")
	}
	http.Handle("/suggestions", handler)

	addr := ":" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
