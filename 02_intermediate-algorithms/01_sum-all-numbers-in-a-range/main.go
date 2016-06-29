/*
	Intermediate Algorithms
	01 Sum All Numbers In A Range

	This program takes two int's and returns
	the sum of all numbers in the range between
	the two.

	Inclusive.
	i.e. 1, 4 = 1, 2, 3, 4 = 10

	Also order of appearance does not matter:
	i.e. 4, 1 = 1, 2, 3, 4 = 10

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// just tests as we dont need data
	fmt.Printf("Sum of range 6 - 1 = %v\n", sumrange(6, 1))
	fmt.Printf("Sum of range 4 - 8 = %v\n", sumrange(4, 8))
	fmt.Printf("Sum of range 0 - 26 = %v\n", sumrange(0, 26))
	// it even works with negative ranges
	fmt.Printf("Sum of range -8 - 2 = %v\n", sumrange(-8, 2))
}

// function to sum numbers within a range
// and return the total as an int
func sumrange(min int, max int) int {
	// init empty total output
	var total int
	// if range was entered max, min: reverse it
	if min > max {
		min, max = max, min
	}
	// until min == max
	for min <= max {
		// add min to total
		total += min
		// increase min
		min++
	}
	// return total
	return total
}
