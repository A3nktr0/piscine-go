package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	temp := 0
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			temp++
		}
	}
	return temp == 0
}
