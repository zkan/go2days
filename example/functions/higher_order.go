package main

import "fmt"

func main() {
	higherOrder(area)
}

func area(dx, dy float64) float64 {
	return dx * dy
}

func higherOrder(firstClass func(float64, float64) float64) {
	fmt.Printf("area of 64x48 = %.0f", firstClass(64, 48))
}
