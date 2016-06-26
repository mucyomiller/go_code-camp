/*
	Basic Algorithms
	11 Slasher Flick

	This program returns the remaining elements of a slice
	of int after slashing of X from the head. Sort of silly
	as it's just a front end for `slice[x:]`.

	NOTE: I went with slice of int as you can't mix types
	in slices in go, so I had to pick one.

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// some data to play with
	numbers1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers2 := []int{453, 22, 657, 11, 105, 44, 73, 28}
	// test the function
	fmt.Printf("Input: %v, slashed at %v\n%v\n", numbers1, 5, slash(numbers1, 5))
	fmt.Printf("Input: %v, slashed at %v\n%v\n", numbers2, 4, slash(numbers2, 4))
}

// function to slash of head of slice
func slash(slc []int, x int) []int {
	// slice the slice at x
	slc = slc[x:]
	// return the slice
	return slc
}
