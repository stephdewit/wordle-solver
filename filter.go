package main

import (
	"strings"
)

func filterWords(words *[]string, word string, result string) []string {
	charsToRemove := []rune{}
	for i, char := range []rune(result) {
		if char == rune('-') {
			charsToRemove = append(charsToRemove, []rune(word)[i])
		}
	}

	filtered := []string{}
	for _, word := range *words {
		contains := false
		for _, char := range charsToRemove {
			if strings.ContainsRune(word, char) {
				contains = true
				break
			}
		}

		if !contains {
			filtered = append(filtered, word)
		}
	}

	return filtered
}
