package main

import (
	"testing"
)

var defaultDataset = []string { "abcde", "cdefg", "hijkl", "mnopl", "qrstq" }

type filterTest struct {
	dataset []string
	proposal, result string
	expected []int
}

var filterTests = []filterTest {
	filterTest{nil, "vwxyz", "-----", []int{0, 1, 2, 3, 4}},
	filterTest{nil, "vwxyz", "-:--!", []int{}},
}

func TestFilterWords(t *testing.T) {
	for _, test := range filterTests {
		dataset := defaultDataset
		if test.dataset != nil {
			dataset = test.dataset
		}

		filtered := filterWords(dataset, test.proposal, test.result)

		if len(filtered) != len(test.expected) {
			t.Errorf("Expected %d words but got %d", len(test.expected), len(filtered))
		}
	}
}
