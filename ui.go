package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func showWords(words []Word, howMuch int) {
	var strings []string
	for i, word := range words {
		strings = append(strings, word.value)

		if i == howMuch-1 {
			break
		}
	}
	fmt.Println(strings)
}

func readInput(label string, characters string, length int) (string, bool, error) {
	fmt.Printf("%s: ", label)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return "", true, nil
		}

		panic(err)
	}

	text = strings.TrimSpace(text)

	if len(text) == 0 {
		return "", false, nil
	}

	match, err := regexp.MatchString(fmt.Sprintf("^[%s]{%d}$", characters, length), text)
	if err != nil {
		panic(err)
	}

	if !match {
		return "", false, errors.New("Invalid input")
	}

	return text, false, nil
}

func readWord(length int) (string, bool) {
	for {
		input, exit, err := readInput("  Word", "a-z", length)
		if err != nil {
			fmt.Println(err)
		} else {
			return input, exit
		}
	}
}

func readResult(length int) (string, bool) {
	for {
		input, exit, err := readInput("Result", "-:!", length)
		if err != nil {
			fmt.Println(err)
		} else if !exit && len(input) == 0 {
			fmt.Println("Missing input")
		} else {
			return input, exit
		}
	}
}
