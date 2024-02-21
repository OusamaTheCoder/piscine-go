package piscine

func BasicAtoi(s string) int {
	reslt := 0
	for _, x := range s {
		d := int(x - '0')
		reslt = reslt*10 + d
	}
	return reslt
}
