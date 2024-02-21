package piscine

func RecursiveFactorial(nb int) int {
	var res int

	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		res = nb * RecursiveFactorial(nb-1)
	}
	return res
}
