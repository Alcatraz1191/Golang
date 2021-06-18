package main

import (
	"fmt"
)

func main() {
	var (
		a string
		//b int
		c float64
	)
	//fmt.Scanln(&a)
	a = "1111"
	fmt.Printf("%T, %v, %v", a, a, c)
}
