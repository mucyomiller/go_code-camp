/*
	Basic Algorithms
	08 Repeat A String

	This program returns a string composed of
	the string(1st input) repeated x(2nd input) times.
	Returning an empty string is x was not positive.

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// define test data
	string1 := "rob"
	string2 := "go"
	// call the function a few times
	fmt.Printf("Repeating %v %v times..\n%v\n", string1, 2, repeat(string1, 2))
	fmt.Printf("Repeating %v %v times..\n%v\n", string2, 5, repeat(string2, 5))
	fmt.Printf("Repeating %v %v times..\n%v\n", string1, -3, repeat(string1, -3))
}

func repeat(str string, x int) string {
	// empty output string
	var output string
	// if x was not negative
	if x > -1 {
		// loop over x
		for x > 0 {
			// concat str to output
			output += str
			// decrease x
			x--
		}
	}
	// return output, empty or otherwise
	return output
}
