package main

import (
	"fmt"
)

func showWords(words *[]string, howMuch int) {
	fmt.Println(getRandomSubArray(words, howMuch))
}
