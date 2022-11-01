package piscine

func FindNextPrime_v2(nb int) int {
	found := false
	next := 2
	if nb > 0 {
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
