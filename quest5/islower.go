package piscine

func IsLower(s string) bool {
	count := 0

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			count++
		}
	}
	return count == len(s)
}
