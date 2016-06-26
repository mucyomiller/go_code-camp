/*
	Basic Algorithms
	10 Chunkey Monkey

	This program splits a slice into x sized
	slices and returns them as a 2D slice.

	Written by: robjloranger
*/
// TODO: Find a way to provide both string and int slice types
// as input to chunky
package main

import "fmt"

func main() {
	// test arrays
	array1 := []string{"a", "b", "c", "d", "e"}
	// array2 := []int{0, 1, 2, 3, 4, 5, 6}
	// run examples
	fmt.Println(chunk(array1, 2))
	fmt.Println(chunk(array1, 4))
	// fmt.Println(chunk(array2, 3))
	// fmt.Println(chunk(array2, 5))
}

// function to return slices, x sized from slc
// as a 2D slice
func chunk(slc []string, x int) [][]string {
	// empty ouput slice of string
	var output [][]string
	for len(slc) > 0 { // while input still contains data
		// create an empty lemon slice 😀
		var lemon []string
		for y := 0; y < x; y++ { // loop over 'y' up to 'x'
			if len(slc)-y > 0 { // if there is data at 'y' in slc
				lemon = append(lemon, slc[y]) // append data to lemon
			}
		} // end 'y' to 'x' loop
		if len(slc) >= x { // if there are more than 'x' data left
			slc = slc[x:] // slice slc at index 'x', non-inclusive
		} else {
			slc = slc[len(slc):] // else slice slc at length to empty it
		}
		output = append(output, lemon) // append lemon to output
	} // end while input contains data loop
	return output // return output of new 2D slice
}
