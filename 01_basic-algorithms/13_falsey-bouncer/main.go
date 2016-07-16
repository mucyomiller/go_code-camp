package main

import "fmt"

func main() {
	data := []interface{}{"hello", 2, true, 0, "", false, "false"}
	fmt.Println(data)
	bounced := bounce(data)
	fmt.Println(bounced)
}

func bounce(a []interface{}) []interface{} {
	var output []interface{}
	for _, v := range a {
		if v == true {
			output = append(output, v)
		}
	}
	return output
}
