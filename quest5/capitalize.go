package piscine

func Capitalize(s string) string {
	str := []rune(s)

	if IsLower(string(str[0])) {
		str[0] -= 32
	}

	for i := 1; i < len(str); i++ {
		if IsAlpha(string(str[i-1])) {
			if IsUpper(string(str[i])) {
				str[i] += 32
			}
		} else if IsLower(string(str[i])) {
			str[i] -= 32
		}
	}
	return string(str)
}
