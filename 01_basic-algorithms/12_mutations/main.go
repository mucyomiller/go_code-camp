/*
	Basic Algorithms
	12 Mutations

	This program checks to see if all letters
	from one string are present in the other.

	Written by: robjloranger
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	// data for us
	first1, first2, first3 := "hello", "hey", "Hello"
	second1, second2 := "Alien", "line"
	// test the function
	fmt.Printf("Does %s contain %s?\n%v\n", first1, first2, mutants(first1, first2))
	fmt.Printf("Does %s contain %s?\n%v\n", first1, first3, mutants(first1, first3))
	fmt.Printf("Does %s contain %s?\n%v\n", second1, second2, mutants(second1, second2))
}

// function to check for matches
func mutants(str1 string, str2 string) bool {
	str1 = strings.ToLower(str1) // convert both inputs to lower
	str2 = strings.ToLower(str2) // case to check for word matches
	// convert second input to slice of runes
	test := []rune(str2)
	// iterate over runes
	for _, r := range test {
		// if rune does not exist in first string, return false
		if strings.IndexRune(str1, r) < 0 {
			return false
		}
	}
	// otherwise return true
	return true
}
