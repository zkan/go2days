package main

import "strconv"

func FizzBuzz(n int) string {
	var s string
	switch {
	case n%15 == 0:
		s = "FizzBuzz"
	case n%5 == 0:
		s = "Buzz"
	case n%3 == 0:
		s = "Fizz"
	default:
		s = strconv.Itoa(n)
	}
	return s
}
