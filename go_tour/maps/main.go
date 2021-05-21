package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	count := 1
	for _, v := range strings.Fields(s){
		m[v] = m[v] + count
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
