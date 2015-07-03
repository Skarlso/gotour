package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	my_slice := make([][]uint8, dy)
	for i := range (my_slice) {
		my_slice[i] = make([]uint8, dx)
		for j := range (my_slice[i]) {
			my_slice[i][j] = uint8(i*j+(i+j)/2)
		}
	}
	
	return my_slice
}

func main() {
	pic.Show(Pic)
}

