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
	sentence1 := "Hello there world"
	sentence2 := "I'm A liTTlE teA Pot"
	sentence3 := "SHORT AND STOUT"
	fmt.Println(titlecase(sentence1))
	fmt.Println(titlecase(sentence2))
	fmt.Println(titlecase(sentence3))
}

func titlecase(s string) string {
	s = strings.ToLower(s)
	return strings.Title(s)
}
