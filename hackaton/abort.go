package piscine

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	var m int
	for i := 0; i <= len(tab)-1; i++ {
		for j := 0; j < len(tab)-1-i; j++ {
			if tab[j] > tab[j+1] {
				tab[j], tab[j+1] = tab[j+1], tab[j]
			}
		}
	}
	m = tab[len(tab)/2]
	return m
}
