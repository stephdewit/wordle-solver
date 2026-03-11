package main

type Guess struct {
	Word   string `json:"w"`
	Result string `json:"r"`
}

type SuggestRequest struct {
	Guesses []Guess `json:"g"`
}

type SuggestResponse struct {
	Suggestions []string `json:"s"`
	Probes      []string `json:"p"`
	Total       int      `json:"t"`
}
