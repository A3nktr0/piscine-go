package piscine

func Rot14(s string) string {
	var str string
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r += 14
			if r > 'z' {
				r = (r - 122) + 96
			}
			str += string(r)
		} else if r >= 'A' && r <= 'Z' {
			r += 14
			if r > 'Z' {
				r = (r - 90) + 64
			}
			str += string(r)
		} else {
			str += string(r)
		}
	}
	return str
}
