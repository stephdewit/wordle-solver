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
	filterTest{nil, "vmxym", "-:---", []int{3}},
	filterTest{nil, "vqqyz", "-::--", []int{4}},
	filterTest{nil, "qwqyz", "!-:--", []int{4}},
	filterTest{nil, "qwxyq", "!---!", []int{4}},
	filterTest{nil, "vooyz", "--!--", []int{3}},
	filterTest{[]string{"qbcdq", "aqcde"}, "vwqqz", "--:--", []int{1}},
	filterTest{[]string{"qbcdq", "aqcde"}, "vwqqz", "--::-", []int{0}},
}

func TestFilterWords(t *testing.T) {
	for _, test := range filterTests {
		dataset := defaultDataset
		if test.dataset != nil {
			dataset = test.dataset
		}

		var words []Word
		for _, str := range dataset {
			words = append(words, Word{value: str})
		}

		filtered := filterWords(words, test.proposal, test.result)

		expectedWords := []Word{}
		for _, i := range test.expected {
			expectedWords = append(expectedWords, Word{value: dataset[i]})
		}

		if len(filtered) != len(test.expected) || !reflect.DeepEqual(filtered, expectedWords) {
			t.Errorf("Expected %d words (%v) but got %d (%v) for %s/%s", len(test.expected), expectedWords, len(filtered), filtered, test.proposal, test.result)
		}
	}
}
