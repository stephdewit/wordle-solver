package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const WordList = "/usr/share/dict/words"

func loadWords(length int) ([]string, error) {
	command := exec.Command("/bin/grep", "-i", fmt.Sprintf("^[a-z]\\{%d\\}$", length), WordList)

	var stdout bytes.Buffer
	command.Stdout = &stdout

	err := command.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to run grep: %w", err)
	}

	return strings.Split(stdout.String(), "\n"), nil
}
