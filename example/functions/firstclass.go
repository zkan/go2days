package main

import "fmt"

func main() {
	var add = func(a, b int) int {
		return a + b
	}

	fmt.Println(add(1, 2))
}
