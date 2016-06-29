/*
	Basic Algorithms
	15 Where Do I Belong

	This program determines the index at which an integer
	should be placed in a slice after being sorted.

	Written by: robjloranger
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// data
	numbers := []float64{10, 20, 30, 40, 50}
	// tests
	// set testXidx to each test case results
	// then display results in formatted print
	test1idx, test1full := ibelong(numbers, 35)
	fmt.Printf("If we sort %v, and want to add %v it should be at index %v.\nLike this: %v\n", numbers, 35, test1idx, test1full)
	test2idx, test2full := ibelong(numbers, 10.5)
	fmt.Printf("If we sort %v, and want to add %v it should be at index %v.\nLike this: %v\n", numbers, 10.5, test2idx, test2full)
}

// function to check where the newbie belongs in our
// group of friends
// takes a slice of float64 and a single float64
func ibelong(friends []float64, newbie float64) (int, []float64) {
	// first add the new float to our slice
	friends = append(friends, newbie)
	// then sort the slice in increasing order
	sort.Float64s(friends)
	// set belongshere to the index of the new float, newbie
	belongshere := sort.SearchFloat64s(friends, newbie)
	// return the index, then the new version of the slice
	return belongshere, friends
}
