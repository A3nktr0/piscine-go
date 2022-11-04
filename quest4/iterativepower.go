package piscine

func IterativePower(nb int, power int) int {
	res := nb

	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return res
	}
	for i := 1; i < power; i++ {
		res *= nb
	}
	return res
}
