package piscine

func StrRev(s string) string {
	str := []rune(s)
	for x, y := 0, len(str)-1; x < y; x, y = x+1, y-1 {
		str[x], str[y] = str[y], str[x]
	}
	return string(str)
}
