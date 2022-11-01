package piscine

func IsNumeric(s string) bool {
	minNum := 48
	maxNum := 57
	count := 0

	for _, v := range s {
		if v >= rune(minNum) && v <= rune(maxNum) {
			count++
		}
	}
	return count == len(s)
}
