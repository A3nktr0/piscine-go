package piscine

func IsPrintable(s string) bool {
	count := 0

	for _, v := range s {
		if v >= 32 {
			count++
		} else {
			count--
		}
	}
	return count == len(s)
}
