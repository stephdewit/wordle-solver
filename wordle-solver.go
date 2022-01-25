package main

import (
	"fmt"
	"log"
)

const Length = 5

func main() {
	words, err := loadWords(Length)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Printf("Found %d %d-letter words\n", len(words), Length)
		showWords(&words, 10)

		break
	}
}
