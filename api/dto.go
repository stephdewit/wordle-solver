package main

import "fmt"

const maxGuesses = 20

type Guess struct {
	Word   string `json:"w"`
	Result string `json:"r"`
}

type SuggestRequest struct {
	Guesses []Guess `json:"g"`
}

func (r SuggestRequest) validate(wordLen int) error {
	if len(r.Guesses) > maxGuesses {
		return fmt.Errorf("too many guesses (max %d)", maxGuesses)
	}
	for _, g := range r.Guesses {
		if len(g.Word) != wordLen {
			return fmt.Errorf("word %q must be %d letters", g.Word, wordLen)
		}
		for _, ch := range g.Word {
			if ch < 'a' || ch > 'z' {
				return fmt.Errorf("word %q contains invalid character", g.Word)
			}
		}
		if len(g.Result) != wordLen {
			return fmt.Errorf("result for %q must be %d characters", g.Word, wordLen)
		}
		for _, ch := range g.Result {
			if ch != '!' && ch != ':' && ch != '-' {
				return fmt.Errorf("result for %q contains invalid character", g.Word)
			}
		}
	}
	return nil
}

type SuggestResponse struct {
	Suggestions []string `json:"s"`
	Probes      []string `json:"p"`
	Total       int      `json:"t"`
}
