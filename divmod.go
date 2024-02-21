package piscine

func DivMod(a int, b int, div *int, mod *int) {
	x := a / b
	*div = x
	*mod = a % b
}
