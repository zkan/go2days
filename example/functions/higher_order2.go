package main

import "fmt"

func main() {
	fmt.Println(higherOrder()(3, 4))
}

func area(dx, dy float64) float64 {
	return dx * dy
}

func higherOrder() (firstClass func(float64, float64) float64) {
	return area
}
