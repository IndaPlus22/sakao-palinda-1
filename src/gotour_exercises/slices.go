package main

import ("golang.org/x/tour/pic")

func Pic(dx, dy int) [][]uint8 {
	
	thing := make([][]uint8, dy)
	for i := range thing {
		thing[i] = make([]uint8, dx)
	}
	

	for i := 0; i <dy; i++ {
		for j := 0; j < dx; j++ {
			thing[i][j] = uint8(i * j)

		}
	}

	return thing
}

func main() {
	pic.Show(Pic)
}
