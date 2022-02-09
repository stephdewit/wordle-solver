package main

type Word struct {
	value string
}

func word(value string) Word {
	return Word{value: value}
}
