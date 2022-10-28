package piscine

func ConvertToNumber2(n rune) int {
	nb := 0
	if n >= '0' && n <= '9' {
		nb = int(n) - 48
	}
	return nb
}

func BasicAtoi2(s string) int {
	nb := 0
	for i, v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		if i != len(s)-1 {

			nb *= 10
			nb += ConvertToNumber2(v) * 10
		} else {
			nb += ConvertToNumber2(v)
		}
	}
	return nb
}
