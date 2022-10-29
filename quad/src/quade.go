package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {

	if x < 0 || y < 0 {
		return
	}

	for h := 0; h < y; h++ {
		for w := 0; w < x; w++ {
			if h == 0 {
				if w == 0 {
					z01.PrintRune('A')
				} else if w == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			} else if h == y-1 {
				if w == 0 {
					z01.PrintRune('C')
				} else if w == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			}
			if h != 0 && h != y-1 {
				if w == 0 || w == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
