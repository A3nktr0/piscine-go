package piscine

func ToLower(s string) string {
	minAlpha := 65
	maxAlpha := 90
	str := []rune(s)
	lowstr := []rune(s)

	for i := range str {
		if str[i] >= rune(minAlpha) && str[i] <= rune(maxAlpha) {
			lowstr[i] = str[i] + 32
		}
	}
	return string(lowstr)
}
