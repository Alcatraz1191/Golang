package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Hello, playground")
	fmt.Println("Hello, playground")
	fmt.Println("Hello, playground")
	fmt.Scanln(&a)
	b = 9 - a
	result := a * 100 + 90 + b
	fmt.Println("Answer is :", result)
}
