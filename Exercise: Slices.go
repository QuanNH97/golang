package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		s[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j%15 == 0:
				s[i][j] = 240
			case j%3 == 0:
				s[i][j] = 120
			case j%5 == 0:
				s[i][j] = 150
			default:
				s[i][j] = 100
			}
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
