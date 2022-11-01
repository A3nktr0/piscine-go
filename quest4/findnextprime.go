package piscine

func FindNextPrime(nb int) int {
	next := 0

	if nb <= 2 {
		return 2
	}

	for i := nb; i > 0; i++ {
		next = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				next++
			}
		}
		if next == 2 {
			return i
		}
	}
	return next
}
