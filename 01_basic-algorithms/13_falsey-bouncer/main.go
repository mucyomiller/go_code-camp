/*
	Basic Algorithms
	13 Falsey Bouncer

	This program takes an 'array' of mixed values,
	actually a slice of empty interface, and returns
	a new slice with only the values that would be
	true in Javascript

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	data := []interface{}{"hello", 2, true, 0, "", false, "false"}
	fmt.Println(data)
	bounced := bounce(data)
	fmt.Println(bounced)
}

// function to create a new slice with only the
// 'true' values
func bounce(a []interface{}) []interface{} {
	var output []interface{}
	for _, v := range a {
		// TODO: try a switch here on type
		// to check for 'false' for different types
		// replicating javascript behaviour
		if v == true {
			output = append(output, v)
		}
	}
	return output
}
