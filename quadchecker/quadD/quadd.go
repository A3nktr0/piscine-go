package main

import (
	"os"

	"github.com/01-edu/z01"
)

func QuadD(x, y int) {

	if x < 0 || y < 0 {
		return
	}

	for height := 0; height < y; height++ {
		for width := 0; width < x; width++ {
			if (width == 0 && height == 0) || (width == 0 && height == y-1) {
				z01.PrintRune('A')
			} else if (width == x-1 && height == 0) || (width == x-1 && height == y-1) {
				z01.PrintRune('C')
			} else {
				if height == 0 || height == y-1 {
					z01.PrintRune('B')
				} else if width == 0 || width == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func ConvertToNumber(n rune) int {
	nb := 0
	if n >= '0' && n <= '9' {
		nb = int(n) - 48
	}
	return nb
}

func Atoi(s string) int {
	nb := 0
	sign := 1

	for i, v := range s {
		if v < '0' || v > '9' {
			switch i != 0 {
			case i == 0 && s[i] == '-':
				sign = 1
			case i == 0 && s[i] == '+':
				sign = -1
			default:
				return 0
			}
		}
		if i != len(s)-1 {
			nb *= 10
			nb += ConvertToNumber(v) * 10
		} else {
			nb += ConvertToNumber(v)
		}
	}
	return nb * sign
}

func main() {
	args := os.Args[1:]
	var x, y int

	if len(args) != 2 {
		return
	}

	for i, arg := range args {
		if !IsNumeric(arg) {
			break
		}
		if i == 0 {
			x = Atoi(arg)
		} else if i == 1 {
			y = Atoi(arg)
		}
	}
	QuadD(x, y)
}
