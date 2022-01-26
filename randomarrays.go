package main

import (
	"math/rand"
	"sort"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func getRandomSubArray(words []string, howMuch int) []string {
	if len(words) <= howMuch {
		return words
	}

	indexes := make(map[int]bool)
	for len(indexes) < howMuch {
		indexes[rng.Int()%len(words)] = true
	}

	var subArray []string
	for index, _ := range indexes {
		subArray = append(subArray, words[index])
	}

	sort.Strings(subArray)

	return subArray
}
