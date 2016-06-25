/*
	Basic Algorithms
	07 Confirm The Ending

	This program checks the end of a string for
	a target. "hello", "lo" would return true
	"hello", "l" would return false

	Written by: robjloranger
*/

package main

import "fmt"

func main() {
	// define test data
	string1, target1 := "Hello World", "ld"
	string2, target2 := "golang go!", "go!"
	// execute function a few times
	fmt.Printf("The string '%v' ends in '%v' == %v\n", string1, target1, confirm(string1, target1))
	fmt.Printf("The string '%v' ends in '%v' == %v\n", string1, target2, confirm(string1, target2))
	fmt.Printf("The string '%v' ends in '%v' == %v\n", string2, target2, confirm(string2, target2))
	fmt.Printf("The string '%v' ends in '%v' == %v\n", string2, target1, confirm(string2, target1))
}

func confirm(str, target) bool {

}
