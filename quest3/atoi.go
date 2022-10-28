package piscine

func ConvertToNumber3(n rune) int {
	nb := 0
	if n >= '0' && n <= '9' {
		nb = int(n) - 48
	}
	return nb
}

func Atoi(s string) int {
	nb := 0
	sign := 1

	for i, v := range s {

		if v < '0' || v > '9' {
			switch i != 0 {
			case i == 0 && s[i] == '-':
				sign = 1
			case i == 0 && s[i] == '+':
				sign = -1
			default:
				return 0
			}
		}

		if i != len(s)-1 {
			nb *= 10
			nb += ConvertToNumber3(v) * 10
		} else {
			nb += ConvertToNumber3(v)
		}
	}
	return nb * sign
}
