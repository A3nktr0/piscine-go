package piscine

func FindNextPrime(nb int) int {
	next := 2
	for i := nb; i < 1000*nb; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return next
}
