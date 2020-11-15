package main

import "fmt"

func main() {
	var i interface{}

	i = 10
	fmt.Printf("type is %T, value is %v\n", i, i)

	var n int

	n = i
}
