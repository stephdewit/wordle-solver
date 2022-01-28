package main

import (
	"reflect"
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

		expectedWords := []string{}
		for _, i := range test.expected {
			expectedWords = append(expectedWords, dataset[i])
		}

		if len(filtered) != len(test.expected) || !reflect.DeepEqual(filtered, expectedWords) {
			t.Errorf("Expected %d words (%v) but got %d (%v) for %s/%s", len(test.expected), expectedWords, len(filtered), filtered, test.proposal, test.result)
		}
	}
}
