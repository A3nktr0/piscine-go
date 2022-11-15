package piscine

func ActiveBits(n int) int {
	var bin []int
	var count int
	for n != 0 {
		bin = append(bin, n%2)
		n /= 2
	}
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == 1 {
			count++
		}
	}
	return count
}
