/*
	Basic Algorithms
	04 Find The Longest Word In A String

	Program returns the length of the longest
	word in a given string

	Written by: robjloranger
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var longest int
	sentence := "The quick brown fox jumped over the lazy elephant"
	words := strings.Split(sentence, " ")
	for _, w := range words {
		if len(w) > longest {
			longest = len(w)
		}
	}
	fmt.Printf("The longest word was %v characters long.\n", longest)
}
