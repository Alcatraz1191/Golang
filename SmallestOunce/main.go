//Write a function that takes in a number of ounces and returns the smallest number of bars
//of soap requred to fulfill the order. There are multiple sizes of soap bars: 2z, 5oz, 10oz

package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2, 5, 10}
	b := minBars(15, s)
	fmt.Println(b)
}

func minBars(ounces int, s []int) int {
	c := len(s) - 1
	barCount := 0
	sort.Ints(s)
	for i := 0; ounces > 0; i++ {
		if ounces >= s[c] {
			ounces = ounces - s[c]
			barCount++
		} else {
			c--
		}
		if c < 0 {
			break
		}
	}
	return barCount
}

func minBarsNoResidue(ounces int, s []int) int {
	sort.Ints(s)

	return 0
}
