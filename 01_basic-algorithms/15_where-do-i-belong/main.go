/*
	Basic Algorithms
	15 Where Do I Belong

	This program determines the index at which an integer
	should be placed in a slice after being sorted.

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// data
	numbers := []float64{10, 20, 30, 40, 50}
	// tests
	test1idx, test1full := ibelong(numbers, 35)
	fmt.Printf("If we sort %v, and want to add %v it should be in index %v.\nLike this: %v\n", numbers, 35, test1idx, test1full)
	test2idx, test2full := ibelong(numbers, 10.5)
	fmt.Printf("If we sort %v, and want to add %v it should be in index %v.\nLike this: %v\n", numbers, 10.5, test2idx, test2full)
}

func ibelong(friends []float64, newbie float64) (int, []float64) {
	return 2, friends
}
