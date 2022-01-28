package main

import (
	"fmt"
)

const Length = 5

func main() {
	words, err := loadWords(Length, false)
	if err != nil {
		panic(err)
	}

	for len(words) > 1 {
		fmt.Printf("Found %d %d-letter words\n", len(words), Length)
		showWords(words, 10)

		word := readWord(Length)

		if len(word) == 0 {
			continue
		}

		result := readResult(Length)

		words = filterWords(words, word, result)
	}

	if len(words) == 0 {
		fmt.Println("No more words")
	} else {
		fmt.Println("Only one word left:", words[0])
	}
}
