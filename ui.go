package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func showWords(words *[]string, howMuch int) {
	fmt.Println(getRandomSubArray(words, howMuch))
}

func readInput(label string, characters string, length int) (string, error) {
	fmt.Printf("%s: ", label)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	text = strings.TrimSpace(text)

	if len(text) == 0 {
		return "", nil
	}

	match, err := regexp.MatchString(fmt.Sprintf("^[%s]{%d}$", characters, length), text)
	if err != nil {
		panic(err)
	}

	if !match {
		return "", errors.New("Invalid input")
	}

	return text, nil
}

func readWord(length int) string {
	for {
		input, err := readInput("  Word", "a-z", length)
		if err != nil {
			fmt.Println(err)
		} else {
			return input
		}
	}
}

func readResult(length int) string {
	for {
		input, err := readInput("Result", "-:!", length)
		if err != nil {
			fmt.Println(err)
		} else if len(input) == 0 {
			fmt.Println("Missing input")
		} else {
			return input
		}
	}
}
