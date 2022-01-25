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

	fmt.Printf("Loaded %d %d-letter words\n", len(words), Length)
}
