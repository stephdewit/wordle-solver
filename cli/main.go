package main

import (
	"fmt"

	"github.com/stephdewit/wordle-solver/solver"
)

const Length = 5

func main() {
	allWords, err := solver.LoadWords(Length, false)
	if err != nil {
		panic(err)
	}

	candidates := allWords
	var guessedWords []string

	done := false
	for len(candidates) > 1 {
		fmt.Printf("Found %d %d-letter words\n", len(candidates), Length)
		fmt.Printf("%10s: ", "Candidates")
		showWords(candidates, 10)

		probes := solver.BestProbes(candidates, allWords, guessedWords)
		if len(probes) > 0 {
			fmt.Printf("%10s: ", "Probes")
			showWords(probes, 10)
		}

		word, exit := readWord(Length)
		if exit {
			done = true
			break
		}

		if len(word) == 0 {
			continue
		}

		result, exit := readResult(Length)
		if exit {
			done = true
			break
		}

		guessedWords = append(guessedWords, word)
		candidates = solver.FilterWords(candidates, word, result)
	}

	if done {
		fmt.Println("\nDone")
	} else if len(candidates) == 0 {
		fmt.Println("No more words")
	} else {
		fmt.Println("Only one word left:", candidates[0].Value)
	}
}
