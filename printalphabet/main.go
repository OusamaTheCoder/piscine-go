package main

import "github.com/01-edu/z01"

func main() {
	for x := 'a'; x <= 'z'; x++ {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}
