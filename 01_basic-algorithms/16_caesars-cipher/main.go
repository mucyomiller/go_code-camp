/*
	Basic Algorithms
	16 Caesars Cipher

	This program will decode a string encoded
	in ROT13 cipher. That is the characters are
	shifted 13 places.
	i.e. 'A' -> 'N', 'B' -> 'O'

	All input should be upper case.


	NOTE: Offset
	The offset mention in the rot13 func is the
	difference in character decimal value beyond 65.
	This offset is used against decimal value 90
	effectively wrapping the upper case alphabet
	between 'A' and 'Z' like a loop.

	Written by: robjloranger
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	//some testing data
	code1 := "SERR PBQR PNZC"
	code2 := "SERR YBIR?"
	code3 := "YRNEA TBYNAT!"
	//print out conversion results
	fmt.Printf("%s reads: %s\n", code1, rot13(code1))
	fmt.Printf("%s reads: %s\n", code2, rot13(code2))
	fmt.Printf("%s reads: %s\n", code3, rot13(code3))
}

// function to convert from ROT13
func rot13(s string) string {
	// init empty output string
	var output string
	// convert input to upper case to comply
	// with requirement
	s = strings.ToUpper(s)
	// loop over input string
	for _, c := range s {
		// if c is one of the first 13
		// chars in the alphabet
		if rune(c) >= 65 && rune(c) < 78 {
			// calculate the offset, see description note
			diff := 13 - (rune(c) - 65)
			// append calculated char to output
			output += string(91 - diff)
		} else if rune(c) >= 78 && rune(c) <= 90 {
			// else if c is the 14th to 26th character
			// offset char decimal value by -13
			output += string(c - 13)
		} else {
			// else append value to string
			// this includes spaces and punctuation
			output += string(c)
		}
	}
	// return new output string
	return output
}
