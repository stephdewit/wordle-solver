package main

import (
	"testing"
)

var dataset = []string { "abcde", "cdefg", "hijkl", "mnopl", "qrstq" }

func TestFilterWordsWithNoCommonLettersAndNegativeResult(t *testing.T) {
	filtered := filterWords(dataset, "vwxyz", "-----")

	if (len(filtered) != 5) {
		t.Errorf("Expected 5 words but got %d", len(filtered))
	}
}

func TestFilterWordsWithNoCommonLettersAndPositiveResult(t *testing.T) {
	filtered := filterWords(dataset, "vwxyz", "-:--!")

	if (len(filtered) != 0) {
		t.Errorf("Expected 0 words but got %d", len(filtered))
	}
}
