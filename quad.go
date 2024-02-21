package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x >= 0 && y >= 0 {
		for a := 0; a < x; a++ {
			if a == 0 {
				z01.PrintRune('/')
			} else if a == x-1 {
				z01.PrintRune('\\')
			} else {
				z01.PrintRune('*')
			}
		}
		z01.PrintRune('\n')

		for b := 0; b < y-2; b++ {
			z01.PrintRune('|')
			for a := 0; a < x-2; b++ {
				z01.PrintRune(' ')
			}
			if x != 1 {
				z01.PrintRune('*')
			}
			z01.PrintRune('\n')
		}

	}
}

func main() {
	QuadA(5, 3)
}
