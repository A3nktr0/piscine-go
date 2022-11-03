package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := range a {
		for _, v := range a[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
