package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const WordList = "/usr/share/dict/words"

func loadWords(length int, includeProperNouns bool) ([]string, error) {
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

	return strings.Split(strings.TrimSuffix(stdout.String(), "\n"), "\n"), nil
}
