package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	for i := 0; i < dy; i++ {
		var row []uint8
		for j := 0; j < dy; j++ {
			row = append(row, uint8(i)*uint8(j))
		}
		result = append(result, row)
	}
	return result
}

func main() {
	pic.Show(Pic)
}
