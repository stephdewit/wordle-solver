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
	{nil, "vwxyz", "-----", []int{0, 1, 2, 3, 4}},
	{nil, "vwxyz", "-:--!", []int{}},
	{nil, "vwxyc", "----:", []int{0, 1}},
	{nil, "vwxyl", "----!", []int{2, 3}},
	{nil, "vaxdz", "-:-!-", []int{0}},
	{nil, "hixyz", "!!---", []int{2}},
	{nil, "vwayz", "-----", []int{1, 2, 3, 4}},
	{nil, "vwqyz", "--:--", []int{4}},
	{nil, "vwxyq", "----!", []int{4}},
	{nil, "vwqyq", "--:-!", []int{4}},
	{nil, "vmxym", "-:---", []int{3}},
	{nil, "vqqyz", "-::--", []int{4}},
	{nil, "qwqyz", "!-:--", []int{4}},
	{nil, "qwxyq", "!---!", []int{4}},
	{nil, "vooyz", "--!--", []int{3}},
	{[]string{"qbcdq", "aqcde"}, "vwqqz", "--:--", []int{1}},
	{[]string{"qbcdq", "aqcde"}, "vwqqz", "--::-", []int{0}},
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
