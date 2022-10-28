package piscine

func ConvertToNumber(n rune) int {
	nb := 0
	if n >= '0' || n <= '9' {
		nb = int(n) - 48
	}
	return nb
}

func BasicAtoi(s string) int {
	nb := 0
	for i, v := range s {
		if i != len(s)-1 {

			nb *= 10
			nb += ConvertToNumber(v) * 10
		} else {
			nb += ConvertToNumber(v)
		}
	}
	return nb
}
