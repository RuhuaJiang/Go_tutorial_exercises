package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for i := 0; i < 100; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %g", float64(e))
}

func main() {
	//fmt.Println(Sqrt(2))
	_, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	}
}
