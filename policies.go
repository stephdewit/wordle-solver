package main

import (
	"strings"
)

type FilterPolicy interface {
	reject(word string, proposal string, result string) bool
}

type ContainsUnknownCharactersPolicy struct{}
type MissesProperlyPlacedCharactersPolicy struct{}
type HasMisplacedCharactersPolicy struct{}
type MissesMisplacedCharactersPolicy struct{}

var filteringPolicies = []FilterPolicy{
	ContainsUnknownCharactersPolicy{},
	MissesProperlyPlacedCharactersPolicy{},
	HasMisplacedCharactersPolicy{},
	MissesMisplacedCharactersPolicy{},
}

func (policy ContainsUnknownCharactersPolicy) reject(word string, proposal string, result string) bool {
	resultChars := []rune(result)
	proposalChars := []rune(proposal)
	wordChars := []rune(word)

	for i, resultChar := range resultChars {
		proposalChar := proposalChars[i]

		if resultChar != rune('-') || !strings.ContainsRune(word, proposalChar) {
			continue
		}

		atMost := 0
		for j, otherResultChar := range resultChars {
			if otherResultChar != rune('-') && proposalChars[j] == proposalChar {
				atMost += 1
			}
		}

		if count(wordChars, proposalChar) > atMost {
			return true
		}
	}

	return false
}

func (policy MissesProperlyPlacedCharactersPolicy) reject(word string, proposal string, result string) bool {
	for i, char := range []rune(result) {
		if char == rune('!') && []rune(word)[i] != []rune(proposal)[i] {
			return true
		}
	}

	return false
}

func (policy HasMisplacedCharactersPolicy) reject(word string, proposal string, result string) bool {
	for i, char := range []rune(result) {
		if char == rune(':') && []rune(word)[i] == []rune(proposal)[i] {
			return true
		}
	}

	return false
}

func (policy MissesMisplacedCharactersPolicy) reject(word string, proposal string, result string) bool {
	for i, char := range []rune(result) {
		if char == rune(':') && !strings.ContainsRune(word, []rune(proposal)[i]) {
			return true
		}
	}

	return false
}

func count(word []rune, char rune) int {
	c := 0

	for _, wordChar := range word {
		if wordChar == char {
			c += 1
		}
	}

	return c
}
