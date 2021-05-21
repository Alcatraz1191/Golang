%package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintln("cannot Sqrt negative number: %v ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := float64(x/2)
	t := 0.0
	if x < 0 {
		return z, ErrNegativeSqrt(x)
	}
	for i := 0; int(100000000*t) != int(100000000*z); i++ {
		t = z
		z -= (z*z - x) / (2*z)
	}
	return z, nil
}

func main() {
	i, err := Sqrt(-2)
	if err !=nil{
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

		
}
