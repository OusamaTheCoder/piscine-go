package main

import "github.com/01-edu/z01"

func main() {
	for x := '0'; x <= '9'; x++ {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}
