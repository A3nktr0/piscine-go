package piscine

func Any(f func(string) bool, a []string) bool {
	res := false
	for _, v := range a {
		if f(v) {
			res = true
		}
	}
	return res
}
