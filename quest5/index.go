package piscine

func Index(s string, toFind string) int {
	sub := s[:len(toFind)]

	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == toFind {
			return i
		}
	}
	return -1
}
