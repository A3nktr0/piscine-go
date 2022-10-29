package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {

	if x < 0 || y < 0 {
		return
	}

	for h := 0; h < y; h++ {
		for w := 0; w < x; w++ {
			if h == 0 || h == y-1 {
				if w == 0 || w == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
			if h != 0 && h != y-1 {
				if w == 0 || w == x-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
