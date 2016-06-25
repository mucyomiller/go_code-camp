/*
	Basic Algorithms
	09 Truncate A String

	Returns a string truncated at the provided length
	ending in '...', the length is inclusive of the dots
	unless the string length is == or < 3.

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
	fmt.Println(truncate(string2, 1))
}

// truncate function, returns a new string
func truncate(str string, x int) string {
	// record input length
	length := len(str)
	// create empty output
	var output string
	// if under 3 chars long exclude dots from total length
	if length <= 3 {
		output = str[0:x] // split string
		output += "..."   // append dots
		// else inlcude dots in total length
	} else {
		output = str[0 : x-3] // split string
		output += "..."       // append dots
	}
	// return new truncated string with dots
	return output
}
