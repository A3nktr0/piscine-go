package piscine

func RecursiveFactorial(nb int) int {
	res := 1

	if nb < 0 || nb >= 21 {
		return 0
	}

	if nb == 1 {
		return res
	} else if nb > 1 {
		res = nb * RecursiveFactorial(nb-1)
	}
	return res
}
