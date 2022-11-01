package piscine

func IsAlpha(s string) bool {
	minMaj := 65
	maxMaj := 90
	minMin := 97
	maxMin := 122
	minNum := 48
	maxNum := 57
	count := 0

	for _, v := range s {
		if (v >= rune(minMin)) && (v <= rune(maxMin)) || (v >= rune(minMaj)) && (v <= rune(maxMaj)) || (v >= rune(minNum)) && (v <= rune(maxNum)) {
			count++
		}
	}
	return count == len(s)
}
