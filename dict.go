package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strings"
)

const WordList = "/usr/share/dict/words"

func loadWords(length int, includeProperNouns bool) ([]Word, error) {
	caseParameter := "--no-ignore-case"
	if includeProperNouns {
		caseParameter = "--ignore-case"
	}

	command := exec.Command("/bin/grep", caseParameter, fmt.Sprintf("^[a-z]\\{%d\\}$", length), WordList)

	var stdout bytes.Buffer
	command.Stdout = &stdout

	err := command.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to run grep: %w", err)
	}

	strings := strings.Split(strings.TrimSuffix(stdout.String(), "\n"), "\n")

	frequencies := getFrequencies(strings)

	var words []Word
	for _, s := range strings {
		words = append(words, word(s, frequencies))
	}

	sort.Slice(words, func(i, j int) bool {
		return words[i].weight > words[j].weight
	})

	return words, nil
}
