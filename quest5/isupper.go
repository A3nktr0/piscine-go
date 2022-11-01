package piscine

func IsUpper(s string) bool {
	count := 0

	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			count++
		}
	}
	return count == len(s)
}
