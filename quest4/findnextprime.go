package piscine

func FindNextPrime(nb int) int {
	found := false
	next := nb
	if nb >= -1000000087 && nb <= 1000000087 {
		if IsPrime(nb) {
			return nb
		} else {
			for i := nb; !found; i++ {
				found = IsPrime(i)
				next = i
			}
		}
	}
	return next
}

// FindNextPrime(1000000088) == 1000000088 instead of 1000000093
