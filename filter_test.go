package main

import (
	"reflect"
	"testing"
)

var defaultDataset = []string{"abcde", "cdefg", "hijkl", "mnopl", "qrstq"}

type filterTest struct {
	dataset          []string
	proposal, result string
	expected         []int
}

var filterTests = []filterTest{
	filterTest{nil, "vwxyz", "-----", []int{0, 1, 2, 3, 4}},
	filterTest{nil, "vwxyz", "-:--!", []int{}},
	filterTest{nil, "vwxyc", "----:", []int{0, 1}},
	filterTest{nil, "vwxyl", "----!", []int{2, 3}},
	filterTest{nil, "vaxdz", "-:-!-", []int{0}},
	filterTest{nil, "hixyz", "!!---", []int{2}},
	filterTest{nil, "vwayz", "-----", []int{1, 2, 3, 4}},
	filterTest{nil, "vwqyz", "--:--", []int{4}},
	filterTest{nil, "vwxyq", "----!", []int{4}},
	filterTest{nil, "vwqyq", "--:-!", []int{4}},
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
