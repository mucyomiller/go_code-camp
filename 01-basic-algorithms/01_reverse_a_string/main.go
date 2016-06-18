/*
	Basic Algorithms
	01 Reverse A String

	This program reverses the runes in a string and
	returns them as a new string

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// assign a couple example strings
	string1 := "robbie"
	string2 := "hello world!"
	// print the returned strings
	fmt.Println(reverse(string1))
	fmt.Println(reverse(string2))
}

// reversing function
func reverse(s string) string {
	// convert the string to a slice of runes
	runes := []rune(s)
	// loop over the slice, stopping at halfway
	for x, j := 0, len(runes)-1; x < len(runes)/2; x++ {
		// swap the beginning with the end
		// move one step inward each iteration
		runes[x], runes[j] = runes[j], runes[x]
		j-- // dont forget to move the end point inward
	}
	// return the string value of the reversed slice
	return string(runes)
}
