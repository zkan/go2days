package main

import (
	"fmt"
	"time"
)

func main() {
	total := 10
	now := time.Now()
	for i := 0; i < total; i++ {
		go printout(i)
	}
	fmt.Println(time.Now().Sub(now))
}

func printout(i int) {
	fmt.Println(i)
}
