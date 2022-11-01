package piscine

func FirstRune(s string) rune {
	str := []rune(s)
	var first rune
	for i, v := range str {
		if i == 0 {
			first = rune(v)
		}
	}
	return first
}
