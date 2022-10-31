package piscine

func IterativeFactorial(nb int) int {
	res := 1

	if nb == 0 {
		return res
	} else if nb < 0 || nb >= 21 {
		return res - 1
	}

	for i := 1; i < nb+1; i++ {
		res *= i
	}
	return res
}
