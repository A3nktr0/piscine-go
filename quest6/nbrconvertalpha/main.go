package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ToUpper(r rune) rune {
	minAlpha := 97
	maxAlpha := 122
	var upper rune

	if r >= rune(minAlpha) && r <= rune(maxAlpha) {
		upper = r - 32
	}
	return upper
}

func ConvertToNbr(n rune) int {
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
			nb += ConvertToNbr(v) * 10
		} else {
			nb += ConvertToNbr(v)
		}
	}
	return nb * sign
}

func ConvertNbrToAlphabet(n int) rune {
	var char rune

	if n >= 1 && n <= 26 {
		char = rune(n + 96)
	} else {
		char = ' '
	}
	return char
}

func main() {
	arguments := os.Args[1:]
	upper := 0

	if len(arguments) < 1 {
		return
	}

	if arguments[0] == "--upper" {
		upper = 1
	}
	if upper == 1 {
		arguments = append(arguments[:0], arguments[1:]...)
	}

	for i := range arguments {
		temp := Atoi(arguments[i])
		letter := ConvertNbrToAlphabet(temp)

		if upper == 1 {
			z01.PrintRune(ToUpper(letter))
		} else {
			z01.PrintRune(letter)
		}
	}
	z01.PrintRune('\n')
}
