package main

import "fmt"

func totalMoney(n int) int {
	sum := 0
	monday := 1
	c := 1
	for i := 1; i <= n ; i++ {
		if (i-1) % 7 == 0 && i != 1 {
			monday++
			sum = sum + monday
			c = monday + 1
		} else {
			sum = sum + c
			c++
		}
	}
	return sum
}

func main() {
	fmt.Println(totalMoney(20))
}