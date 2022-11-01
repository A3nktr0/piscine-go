package piscine

func ToUpper(s string) string {
	minAlpha := 97
	maxAlpha := 122
	str := []rune(s)
	upstr := []rune(s)

	for i := range str {
		if str[i] >= rune(minAlpha) && str[i] <= rune(maxAlpha) {
			upstr[i] = str[i] - 32
		}
	}
	return string(upstr)
}
