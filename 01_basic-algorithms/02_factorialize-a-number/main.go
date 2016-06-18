/*
	Basic Algorithms
	02 Factorialize A Number

	This program takes an integer and factorializes it.
	i.e. 5 = 5 * 4 * 3 * 2 * 1 = 120

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Println("Factorialize A Number")
	fmt.Println("---------------------")
	fmt.Printf("Please enter a number to factorialize,\ninput: ")
	fmt.Scan(&num)
	for x := num - 1; x > 0; x-- {
		num *= x
	}
	fmt.Printf("%v is the total of your factorial\n", num)
}
