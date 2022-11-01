package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	var c rune
	if n < 0 || n > len(str) {
		return 0
	} else {
		if n >= 1 {
			c = str[n-1]
		}
	}
	return c
}
