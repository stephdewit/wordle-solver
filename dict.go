package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const defaultWordList = "/usr/share/dict/words"

func resolveWordListPath() string {
	if path := os.Getenv("WORDLIST"); path != "" {
		return path
	}
	return defaultWordList
}

func loadWords(length int, includeProperNouns bool) ([]Word, error) {
	file, err := os.Open(resolveWordListPath())
	if err != nil {
		return nil, fmt.Errorf("failed to open word list: %w", err)
	}
	defer file.Close()

	var strs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w := scanner.Text()
		if isValidWord(w, length, includeProperNouns) {
			strs = append(strs, w)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read word list: %w", err)
	}

	frequencies := getFrequencies(strs)

	words := make([]Word, len(strs))
	for i, s := range strs {
		words[i] = word(s, frequencies)
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].weight > words[j].weight
	})

	return words, nil
}

func isValidWord(s string, length int, includeProperNouns bool) bool {
	if len(s) != length {
		return false
	}
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			continue
		}
		if includeProperNouns && ch >= 'A' && ch <= 'Z' {
			continue
		}
		return false
	}
	return true
}
