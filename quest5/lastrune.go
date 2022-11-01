package piscine

func LastRune(s string) rune {
	str := []rune(s)
	var last rune
	for i, v := range str {
		if i == len(str)-1 {
			last = rune(v)
		}
	}
	return last
}
