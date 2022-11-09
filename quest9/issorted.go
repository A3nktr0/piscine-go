package piscine

func F(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	x, y := true, true

	if len(a) <= 1 {
		return true
	}
	for i := 1; i < len(a); i++ {
		if !(F(a[i-1], a[i]) <= 0) {
			x = false
		} else if !(F(a[i-1], a[i]) >= 0) {
			y = false
		} else {
			return true
		}
	}
	return x != y
}
