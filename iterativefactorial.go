package piscine

func IterativeFactorial(nb int) int {
	var res int = 1

	if nb < 0 || nb > 20 {
		return 0
	} else if nb >= 0 && nb <= 20 {
		for a := 1; a <= nb; a++ {
			res = res * a
		}
	} else {
		return 1
	}
	return res
}
