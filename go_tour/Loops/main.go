package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x/2)
	t := 0.0
	for i := 0; int(100000000*t) != int(100000000*z); i++ {
		t = z
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	x := float64(5)
	fmt.Println(Sqrt(x))
	fmt.Println("----")
	fmt.Println(math.Sqrt(x))
}
