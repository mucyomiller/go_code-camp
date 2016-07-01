/*
	Intermediate Algorithms
	02 Diff Two Arrays

	This program returns the symetric difference
	of two slices. A new slice with values found
	in one slice or the other but not the both.

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// data
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{2, 4, 6, 8, 10}
	slice3 := []int{10, 20, 30, 40}
	// test
	fmt.Printf("The symetric difference between %v and %v is:\n%v\n", slice1, slice2, diff(slice1, slice2))
	fmt.Printf("The symetric difference between %v and %v is:\n%v\n", slice2, slice3, diff(slice2, slice3))
	fmt.Printf("The symetric difference between %v and %v is:\n%v\n", slice1, slice3, diff(slice1, slice3))
}

// function to compare two slices
func diff(one []int, two []int) []int {
	// init empty output slice
	var output []int
	// loop over slices checking if
	// value exists in the other slice
	// if it does not exist, append it to
	// our output
	for _, o := range one { // slice one
		if !exists(o, two) {
			output = append(output, o)
		}
	}
	for _, t := range two { // slice two
		if !exists(t, one) {
			output = append(output, t)
		}
	}
	// return output
	return output
}

// function to check if value exists in a given slice
func exists(x int, list []int) bool {
	// loops over slice
	for _, v := range list {
		if v == x {
			// if value matches input x return true
			return true
		}
	}
	// otherwise return false as it doesn't exist
	return false
}
