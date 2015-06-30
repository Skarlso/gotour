package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	close_enough := false
	previous_calc := 0.1
	iter_count := 0
	for close_enough != true {
		z = z - ((z * z - x) / (2 * z))
		if previous_calc == z || math.Abs(z - previous_calc) < 0.0000001 {
			close_enough = true
		}
		previous_calc = z
		iter_count++
	}
	fmt.Println(iter_count)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
