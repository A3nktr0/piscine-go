package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		count := 0
		for j := 0; j < len(a); j++ {
			if i != j && a[i] == a[j] {
				count++
			}
		}
		if count == 0 || count%2 == 0 {
			return a[i]
		}
	}
	return -1
}
