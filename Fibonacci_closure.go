package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	index := 0
	last_two := 0
	last_one := 0
	return func() int {
		if index == 0 {
			index++
			return 0
		} else if index == 1 {
			index++
			last_two = 0
			last_one = 1
			return 1
		} else {
			index++
			ret := last_one + last_two
			last_two = last_one
			last_one = ret
			return ret

		}

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
