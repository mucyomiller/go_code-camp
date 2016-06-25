/*
	Basic Algorithms
	09 Truncate A String

	Returns a string truncated at the provided length
	ending in '...', the length is inclusive of the dots
	unles the string length is == or < 3.

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// some test data
	string1 := "The quick brown fox"
	string2 := "hi"
	// call truncate tests
	fmt.Println(truncate(string1, 8))
	fmt.Println(truncate(string1, 2))
}

// truncate function, returns a new string
func truncate(str string, x int) string {
	// stuff goes here
}
