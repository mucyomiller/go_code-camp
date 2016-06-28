/*
	Basic Algorithms
	14 Seek And Destroy

	This program removes one or more int's from a slice
	of int.

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// data
	data := []int{4, 5, 2, 4, 2, 1, 5, 7, 8, 9, 7, 6}
	targets := []int{2, 4, 9}
	// tests
	// just setup, not functioning
	fmt.Printf("Data: %v\nDestroy: %v\nResults: %v\n", data, targets, destroy(data, targets))
}

// destructive function, keeps values not on the target list
func destroy(input []int, remove []int) []int {
	// new output slice of int
	var output []int
	// loop over input
	for _, v := range input {
		// if it doesn't exist on the target list
		// append it to the output
		if exists(remove, v) == false {
			output = append(output, v)
		}
	}
	// return output
	return output
}

// function to see if a value matches the target list
func exists(targets []int, value int) bool {
	// loop over the targets
	for _, v := range targets {
		// if the value matches a target, return true
		if v == value {
			return true
		}
	}
	// otherwise return false as it is not present
	return false
}
