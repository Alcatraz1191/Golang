package main

import (
	"golang.org/x/tour/pic";
	"math"

)


func Pic(dx, dy int) [][]uint8 {
	s:= make([][]uint8, dy)
	for v:= range s {
		s[v] = make([]uint8, dx)
		for u:= range s[v]{
			s[v][u] = uint8(math.Sqrt(float64(v*v+u*u)))
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}