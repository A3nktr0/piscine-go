package piscine

import "github.com/01-edu/z01"

func SplitNbr(nb int) []int {
	var tab []int

	for nb != 0 {
		tab = append(tab, nb%10)
		nb /= 10
	}
	return tab
}

func SortNbr(tab []int) int {
	nb := 0

	for i := range tab {
		if tab[i] == 0 {
			z01.PrintRune('0')
		}
	}

	for i := 0; i < len(tab)-1; i++ {
		for j := 0; j < len(tab)-i-1; j++ {
			if tab[j] > tab[j+1] {
				tab[j], tab[j+1] = tab[j+1], tab[j]
			}
		}
	}
	for i := range tab {
		if i != len(tab)-1 {
			nb *= 10
			nb += tab[i] * 10
		} else {
			nb += tab[i]
		}
	}
	return nb
}

func PrintNbrInOrder(n int) {
	nb := '0'

	if n > 0 {
		n = SortNbr(SplitNbr(n))
	}

	if n > 9 {
		PrintNbrInOrder(n / 10)
		n %= 10
	}

	if n >= 0 && n <= 9 {
		for i := 0; i < n; i++ {
			nb++
		}
	}
	z01.PrintRune(nb)
}
