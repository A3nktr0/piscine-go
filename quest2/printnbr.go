package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	nb := '0'

	if n == -9223372036854775808 {
		s := "-9223372036854775808"
		for _, r := range s[:] {
			z01.PrintRune(r)
		}
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}

	if n > 9 {
		PrintNbr(n / 10)
		n %= 10
	}

	if n > 0 && n <= 9 {
		for i := 1; i <= n; i++ {
			nb++
		}
	}
	z01.PrintRune(nb)
}
