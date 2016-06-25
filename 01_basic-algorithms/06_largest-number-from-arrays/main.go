/*
	Basic Algorithms
	06 Return Largest Numbers in Arrays

	Program takes any number of arrays or slices and
	returns a new slice with each of the
	largest numbers found in ALL provided arrays.

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

// function to return slice of largest numbers
// variadic input, output slice of int
func largest(arrays ...[]int) []int {
	// create empty array length of input
	output := make([]int, len(arrays))
	// loop over input
	for i, arr := range arrays {
		// new empty highest number
		var highest int
		// loop over current array/slice of numbers
		for x := 0; x < len(arr)-1; x++ {
			// keep highest up to date
			if arr[x] > highest {
				highest = arr[x]
			}
		}
		// add highest to ouput
		output[i] = highest
	}
	// return output
	return output
}
