/*
	Basic Algorithms
	03 Check For Palindromes

	This program checks if a string is equal to
	itself when reversed, disregarding case or whitespace

	Written by: robjloranger
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Hello World"
	string2 := "racecar"
	string3 := "raCe Car"
	fmt.Printf("String1: %s, proved %v\n", string1, palindrome(string1))
	fmt.Printf("String2: %s, proved %v\n", string2, palindrome(string2))
	fmt.Printf("String3: %s, proved %v\n", string3, palindrome(string3))
}

// function to check if string is spelled the same
// forwards as backwards. returning true if it is
func palindrome(s string) bool {
	// remove capitalization
	s = strings.ToLower(s)
	// remove whitespace
	s = strings.Map(func(rn rune) rune {
		if rn != 32 {
			return rn
		}
		return -1
	}, s)
	// create slice copy
	rv := []rune(s)
	// reverse string
	for x, j := 0, len(rv)-1; x < len(rv)/2; x, j = x+1, j-1 {
		rv[x], rv[j] = rv[j], rv[x]
	}
	// check if reverse is equal to original
	if s == string(rv) {
		return true
	}
	return false
}
