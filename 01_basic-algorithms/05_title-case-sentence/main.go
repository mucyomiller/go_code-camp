/*
	Basic Algorithms
	05 Title Case A Sentence

	Program returns a sentence, title cased

	Written by: robjloranger
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// data to test
	sentence1 := "Hello there world"
	sentence2 := "I'm A liTTlE teA Pot"
	sentence3 := "SHORT AND STOUT"
	// call titlecase on our test data
	fmt.Println(titlecase(sentence1))
	fmt.Println(titlecase(sentence2))
	fmt.Println(titlecase(sentence3))
}

// function to title case a string by words
// i.e. each word title cased
func titlecase(s string) string {
	// start by lower casing the whole string
	s = strings.ToLower(s)
	// set the previous to a space " " so the
	// first word gets title cased
	prev := " "
	// init empty output
	var output string
	// loop over input string
	for _, c := range s {
		// if previous character was a space and this
		// character is alpha
		if prev == " " && rune(c) >= 97 && rune(c) <= 122 {
			// convert to upper case
			output += string(rune(c) - 32)
		} else {
			// else just append to output
			output += string(c)
		}
		// set prev to the last character iterated over
		prev = string(c)
	}
	// return new output string
	return output
}
