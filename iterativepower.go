package piscine

func IterativePower(nb int, power int) int {
	var res int = 1
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		res *= nb
	}
	return res
}
