package main

import (
	"strings"
)

func filterWords(words *[]string, proposal string, result string) []string {
	charsToRemove := []rune{}
	for i, char := range []rune(result) {
		if char == rune('-') {
			charsToRemove = append(charsToRemove, []rune(proposal)[i])
		}
	}

	filtered := []string{}
	for _, word := range *words {
		if containsUnknownCharacters(word, proposal, result) {
			continue
		}

		if missesProperlyPlacedCharacters(word, proposal, result) {
			continue
		}

		if hasMisplacedCharacters(word, proposal, result) {
			continue
		}

		filtered = append(filtered, word)
	}

	return filtered
}

func containsUnknownCharacters(word string, proposal string, result string) bool {
	for i, char := range []rune(result) {
		if char == rune('-') && strings.ContainsRune(word, []rune(proposal)[i]) {
			return true
		}
	}

	return false
}

func missesProperlyPlacedCharacters(word string, proposal string, result string) bool {
	for i, char := range []rune(result) {
		if char == rune('!') && []rune(word)[i] != []rune(proposal)[i] {
			return true
		}
	}

	return false
}

func hasMisplacedCharacters(word string, proposal string, result string) bool {
	for i, char := range []rune(result) {
		if char == rune(':') && []rune(word)[i] == []rune(proposal)[i] {
			return true
		}
	}

	return false
}
