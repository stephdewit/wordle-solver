package main

import (
	"fmt"

	"github.com/stephdewit/wordle-solver/solver"
)

const Length = 5

func main() {
	words, err := solver.LoadWords(Length, false)
	if err != nil {
		panic(err)
	}

	done := false
	for len(words) > 1 {
		fmt.Printf("Found %d %d-letter words\n", len(words), Length)
		showWords(words, 10)

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

		words = solver.FilterWords(words, word, result)
	}

	if done {
		fmt.Println("\nDone")
	} else if len(words) == 0 {
		fmt.Println("No more words")
	} else {
		fmt.Println("Only one word left:", words[0].Value)
	}
}
