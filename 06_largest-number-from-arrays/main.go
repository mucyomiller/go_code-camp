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
	array1 := []int{2, 30, 86, 12}
	array2 := []int{0, 31, 126, 76, 45, 22, 90, 86}
	fmt.Println(largest(array1))
}

func largest(arrays ...[]int) int {
	for arr := range arrays {
		fmt.Println(arr)
	}
	return 1
}
