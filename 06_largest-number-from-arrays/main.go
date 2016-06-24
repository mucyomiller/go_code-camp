/*
	Basic Algorithms
	06 Return Largest Numbers in Arrays

	Program takes any number of arrays and
	returns a new array with each of the
	largest numbers found in any given array.

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// define some test data
	array1 := []int{2, 30, 86, 12}
	array2 := []int{0, 31, 126, 76, 45, 22, 90, 86}
	// execute our function a few ways
	fmt.Println(largest(array1))
	fmt.Println(largest(array1, array2))
	fmt.Println(largest(array1, array2, array1))
}

// function to return array of largest numbers
func largest(arrays ...[]int) []int {
	// variadic input
	// create empty array length of input
	output := make([]int, len(arrays))
	// loop over input
	for arr := range arrays {
		// new empty highest number
		var highest int
		// loop over this array of numbers
		// BUG: Error, invalid input for len.
		// arr type int.. what?
		for x := 0; x < len(arr)-1; x++ {
			// keep highest up to date
			if x > highest {
				highest = x
			}
		}
		// add highest to array
	}
	// return new array
	return output
}
